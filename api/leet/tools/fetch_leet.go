package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const LeetCodeAPI = "https://leetcode.com/graphql"

type SubmitStats struct {
	AcSubmissionNum []struct {
		Difficulty  string `json:"difficulty"`
		Count       int    `json:"count"`
		Submissions int    `json:"submissions"`
	} `json:"acSubmissionNum"`
	TotalSubmissionNum []struct {
		Difficulty  string `json:"difficulty"`
		Count       int    `json:"count"`
		Submissions int    `json:"submissions"`
	} `json:"totalSubmissionNum"`
}

type StreakStats struct {
	TotalSubmissions int `json:"totalSubmissions"`
	CurrentStreak    int `json:"currentStreak"`
	LongestStreak    int `json:"longestStreak"`
}

type UserData struct {
	Data struct {
		AllQuestionsCount []struct {
			Difficulty string `json:"difficulty"`
			Count      int    `json:"count"`
		} `json:"allQuestionsCount"`
		MatchedUser struct {
			Username           string     `json:"username"`
			FirstName          string     `json:"firstName"`
			LastName           string     `json:"lastName"`
			Contributions      struct{ Points int } `json:"contributions"`
			Profile            struct {
				Reputation int    `json:"reputation"`
				Ranking    int    `json:"ranking"`
				UserAvatar string `json:"userAvatar"`
			} `json:"profile"`
			SubmissionCalendar string      `json:"submissionCalendar"`
			SubmitStats        SubmitStats `json:"submitStats"`
			Streak             StreakStats `json:"streak"` // 👈 novo campo
		} `json:"matchedUser"`
		RecentSubmissionList []struct {
			Title         string `json:"title"`
			TitleSlug     string `json:"titleSlug"`
			Timestamp     string `json:"timestamp"`
			StatusDisplay string `json:"statusDisplay"`
			Lang          string `json:"lang"`
		} `json:"recentSubmissionList"`
	} `json:"data"`
}

func GetUserData(username string) (*UserData, error) {
	query := fmt.Sprintf(`{
		allQuestionsCount {
			difficulty
			count
		}
		matchedUser(username: "%s") {
			username
			firstName
			lastName
			contributions {
				points
			}
			profile {
				reputation
				ranking
				userAvatar
			}
			submissionCalendar
			submitStats {
				acSubmissionNum {
					difficulty
					count
					submissions
				}
				totalSubmissionNum {
					difficulty
					count
					submissions
				}
			}
		}
		recentSubmissionList(username: "%s") {
			title
			titleSlug
			timestamp
			statusDisplay
			lang
		}
	}`, username, username)

	reqBody := map[string]interface{}{
		"query": query,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(LeetCodeAPI, "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data UserData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	// 📌 Processa a calendar string para calcular streaks
	streakStats := calculateStreaks(data.Data.MatchedUser.SubmissionCalendar)
	data.Data.MatchedUser.Streak = streakStats

	return &data, nil
}





func calculateStreaks(calendar string) StreakStats {
	var stats StreakStats
	if calendar == "" {
		return stats
	}
	//fmt.Println("Calendar:", calendar)

	var submissionMap map[string]int
	if err := json.Unmarshal([]byte(calendar), &submissionMap); err != nil {
		return stats
	}
	//fmt.Println("Submissions:", submissionMap)

	dateMap := make(map[string]bool)
	for tsStr := range submissionMap {
		tsInt, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			continue
		}
		day := time.Unix(tsInt, 0).UTC().Format("2006-01-02")
		dateMap[day] = true
		stats.TotalSubmissions++
	}
	//fmt.Println("Total de submissões:", stats.TotalSubmissions)
	// Pega todas as datas e ordena
	var dates []string
	for d := range dateMap {
		dates = append(dates, d)
	}
	//fmt.Println("Datas:", dates)
	sort.Strings(dates)

	// Calcula streaks
	var maxStreak, currentStreak int
	var lastDay time.Time

	for i, d := range dates {
		currentDay, err := time.Parse("2006-01-02", d)
		if err != nil {
			continue
		}
		// Tolerância de 1h para considerar dias consecutivos
		if i == 0 || currentDay.Sub(lastDay).Hours() <= 25 {
			currentStreak++
		} else {
			currentStreak = 1
		}
		if currentStreak > maxStreak {
			maxStreak = currentStreak
		}
		lastDay = currentDay
	}

	// Verifica streak atual (até hoje)
	today := time.Now().UTC().Format("2006-01-02")
	if _, ok := dateMap[today]; !ok {
		currentStreak = 0
	}

	stats.CurrentStreak = currentStreak
	stats.LongestStreak = maxStreak

	return stats
}
