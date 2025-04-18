package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type DuolingoResponse struct {
	Users []User `json:"users"`
}


type LanguageXP struct {
	LanguageID string  `json:"language_id"`
	Percentage float64 `json:"percentage"`
}

type User struct {
	Username     string   `json:"username"`
	Name         string   `json:"name"`
	FirstName    string   `json:"firstName"`
	LastName     string   `json:"lastName"`
	Bio          string   `json:"bio"`
	Picture      string   `json:"picture"`
	CreationDate int64    `json:"creationDate"`
	Streak       int      `json:"streak"`
	TotalXP      int      `json:"totalXp"`
	Courses      []Course `json:"courses"`
	StreakData   struct {
		CurrentStreak struct {
			StartDate string `json:"startDate"`
			Length   int    `json:"length"`
			EndDate  string `json:"endDate"`
		} `json:"currentStreak"`
	} `json:"streakData"`
	XPByLanguage []LanguageXP `json:"xp_by_language"` // Added field

}

type Course struct {
	Title           string `json:"title"`
	LearningLanguage string `json:"learningLanguage"`
	FromLanguage     string `json:"fromLanguage"`
	XP              int    `json:"xp"`
	Crowns          int    `json:"crowns"`
	ID              string `json:"id"`
}

func FetchDuolingoUser(user string) (User, error) {
	url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users?username=%s&fields=streak,streakData%%7BcurrentStreak,previousStreak%%7D%%7D", user)
	resp, err := http.Get(url)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	var data DuolingoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return User{}, err
	}

	if len(data.Users) == 0 {
		return User{}, fmt.Errorf("usuário não encontrado")
	}

	userData := data.Users[0]

	// Usa a função separada
	userData.XPByLanguage = CalculateXPByLanguage(userData.Courses)

	return userData, nil
}



func CalculateXPByLanguage(courses []Course) []LanguageXP {
	totalXP := 0
	languageXP := make(map[string]int)

	for _, course := range courses {
		languageXP[course.Title] += course.XP
		totalXP += course.XP
	}

	var xpPercentages []LanguageXP
	if totalXP > 0 {
		for lang, xp := range languageXP {
			percentage := (float64(xp) / float64(totalXP)) * 100
			xpPercentages = append(xpPercentages, LanguageXP{
				LanguageID: lang,
				Percentage: percentage,
			})
		}
	}

	// Ordena do maior para o menor
	sort.Slice(xpPercentages, func(i, j int) bool {
		return xpPercentages[i].Percentage > xpPercentages[j].Percentage
	})

	return xpPercentages
}
