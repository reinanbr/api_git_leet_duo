package handler

import (
	"api_git_leet_duo/api/git/tools/contribuitions"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func GitContrib(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	startingYear := 2015
	graphs, err := contribuitions.GetContributionGraphs(username, startingYear)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving graphsContribuitions: %v", err), http.StatusInternalServerError)
		return
	}

	sortYears := []int{}
	for year := range graphs {
		sortYears = append(sortYears, year)
	}
	sort.Ints(sortYears)

	maxStreak, currentStreak := contribuitions.GetContributionStreaks(graphs)

	// Monta a resposta JSON
	response := make(map[string]interface{})
	response["user"] = username
	response["contributions"] = graphs
	response["streak"] = map[string]interface{}{"max_streak": maxStreak, "current_streak": currentStreak}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
