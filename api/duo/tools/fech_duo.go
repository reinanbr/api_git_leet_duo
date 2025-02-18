package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DuolingoResponse struct {
	Users []User `json:"users"`
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
	url := fmt.Sprintf("https://www.duolingo.com/2017-06-30/users?username=%s&fields=streak,streakData%7BcurrentStreak,previousStreak%7D%7D", user)
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

	return data.Users[0], nil
}