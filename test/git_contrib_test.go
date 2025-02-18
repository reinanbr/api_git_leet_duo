package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const githubAPI = "https://api.github.com/graphql"

type ContributionGraph struct {
	TotalContributions int `json:"totalContributions"`
	Weeks             []struct {
		ContributionDays []struct {
			ContributionCount int    `json:"contributionCount"`
			Date              string `json:"date"`
		} `json:"contributionDays"`
	} `json:"weeks"`
}

type GraphQLResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar ContributionGraph `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}

func getGitHubTokens() []string {
	tokens := os.Getenv("GITHUB_TOKENS")
    //fmt.print("token:"+tokens);
	if tokens == "" {
		return []string{}
	}
	return strings.Split(tokens, ",")
}

func getGitHubToken() string {
	tokens := getGitHubTokens()
	if len(tokens) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	return tokens[rand.Intn(len(tokens))]
}

func fetchContributionGraph(username string, year int) (*ContributionGraph, error) {
	startDate := fmt.Sprintf("%d-01-01T00:00:00Z", year)
	endDate := fmt.Sprintf("%d-12-31T23:59:59Z", year)
	query := fmt.Sprintf(`{"query": "query { user(login: \"%s\") { contributionsCollection(from: \"%s\", to: \"%s\") { contributionCalendar { totalContributions weeks { contributionDays { contributionCount date } } } } } }" }`, username, startDate, endDate)
	token := getGitHubToken()
	if token == "" {
		return nil, fmt.Errorf("no GitHub token available")
	}

	req, err := http.NewRequest("POST", githubAPI, strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "GitHub-Contributions-API")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gqlResponse GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&gqlResponse); err != nil {
		return nil, err
	}

	return &gqlResponse.Data.User.ContributionsCollection.ContributionCalendar, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' query parameter", http.StatusBadRequest)
		return
	}

	year := time.Now().Year()
	graph, err := fetchContributionGraph(username, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph)
}

func main() {
	http.HandleFunc("/contributions", handler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
