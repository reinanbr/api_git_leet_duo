package tools

import (
	"encoding/json"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const LeetCodeAPI = "https://leetcode.com/graphql"

type SubmitStats struct {
	AcSubmissionNum []struct {
		Difficulty string `json:"difficulty"`
		Count      int    `json:"count"`
		Submissions int   `json:"submissions"`
	} `json:"acSubmissionNum"`
	TotalSubmissionNum []struct {
		Difficulty string `json:"difficulty"`
		Count      int    `json:"count"`
		Submissions int   `json:"submissions"`
	} `json:"totalSubmissionNum"`
}

type UserData struct {
	Data struct {
		AllQuestionsCount []struct {
			Difficulty string `json:"difficulty"`
			Count      int    `json:"count"`
		} `json:"allQuestionsCount"`
		MatchedUser struct {
			Username          string     `json:"username"`
			FirstName         string     `json:"firstName"`
			LastName          string     `json:"lastName"`
			Contributions     struct{ Points int } `json:"contributions"`
			Profile           struct {
				Reputation int    `json:"reputation"`
				Ranking    int    `json:"ranking"`
				UserAvatar string `json:"userAvatar"`
			} `json:"profile"`
			SubmissionCalendar string      `json:"submissionCalendar"`
			SubmitStats        SubmitStats `json:"submitStats"`
		} `json:"matchedUser"`
		RecentSubmissionList []struct {
			Title       string `json:"title"`
			TitleSlug   string `json:"titleSlug"`
			Timestamp   string `json:"timestamp"`
			StatusDisplay string `json:"statusDisplay"`
			Lang        string `json:"lang"`
		} `json:"recentSubmissionList"`
	} `json:"data"`
}

func GetUserData(username string) (*UserData, error) {
	// Construindo a consulta GraphQL com o username
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

	// Criando o corpo da requisição em JSON
	reqBody := map[string]interface{}{
		"query": query,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Fazendo a requisição POST
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

	return &data, nil
}
