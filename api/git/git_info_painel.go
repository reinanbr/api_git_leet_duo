package handler

import (
	"api_git_leet_duo/api/git/tools/contribuitions"
	"api_git_leet_duo/api/git/tools/languages"
	"api_git_leet_duo/api/git/tools/user"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// GitPainel - HTTP handler for retrieving GitHub user data.
func GitPainel(w http.ResponseWriter, r *http.Request) {
	// Extract the "user" parameter from the request URL
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	// Initialize response map
	response := make(map[string]interface{})

	// Fetch GitHub data and populate the response map
	if err := addGitInfo(response, username); err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving GitHub data: %v", err), http.StatusInternalServerError)
		return
	}

	// Return response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// addGitInfo - Retrieves and adds GitHub user information to the response map
func addGitInfo(response map[string]interface{}, username string) error {
	// Fetch general user information
	userInfo, err := user.FetchUserData(username)
	if err != nil {
		return fmt.Errorf("error retrieving user info: %w", err)
	}
	response["user"] = userInfo
	
	// Fetch user contribution history starting from 2015
	startingYear := 2015
	graphs, err := contribuitions.GetContributionGraphs(username, startingYear)
	if err != nil {
		return fmt.Errorf("error retrieving contribution info: %w", err)
	}

	// Sort contribution years in ascending order
	sortYears := []int{}
	for year := range graphs {
		sortYears = append(sortYears, year)
	}
	sort.Ints(sortYears)

	// Calculate contribution streaks (max and current)
	maxStreak, currentStreak := contribuitions.GetContributionStreaks(graphs)
	response["streak"] = map[string]interface{}{
		"max_streak":    maxStreak,
		"current_streak": currentStreak,
	}
	
	//total contributions
	
	totalContributions := contribuitions.GetTotalContributions(graphs)
	//total contributions by year
	totalContributionsByYear := contribuitions.GetContributionsByYear(graphs)

	response["contribuitions"] = map[string]interface{}{
		"total": totalContributions,
		"by_year": totalContributionsByYear,
	}

	// Fetch detailed language usage statistics
	langs, err := languages.FetchUserLangsFull(username)
	if err != nil {
		return fmt.Errorf("error retrieving language info: %w", err)
	}

	// Calculate language usage percentage
	langsPercent, totalSize := languages.CalculateLanguagePercentage(langs)
	response["langs"] = map[string]interface{}{
		"total_size":   totalSize,
		"lang_percent": langsPercent,
	}

	// Fetch basic repository language information
	langsLite, err := languages.FetchUserLite(username)
	if err != nil {
		return fmt.Errorf("error retrieving language lite info: %w", err)
	}
	response["repos"] = langsLite


	return nil
}
