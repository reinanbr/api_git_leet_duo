package handler

import (
	"api_git_leet_duo/api/git/tools/contribuitions"
	"api_git_leet_duo/api/git/tools/user"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func Git(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	startingYear := 2015
	graphs, err := contribuitions.GetContributionGraphs(username, startingYear)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving graphs: %v", err), http.StatusInternalServerError)
		return
	}

	userInfo, errUser := user.FetchUserData(username)
	if errUser != nil {
		http.Error(w, fmt.Sprintf("Error retrieving graphsUser: %v", errUser), http.StatusInternalServerError)
		return
	}


	sortYears := []int{}
	for year := range graphs {
		sortYears = append(sortYears, year)
	}
	sort.Ints(sortYears)

	// Monta a resposta JSON
	response := make(map[string]interface{})
	response["user"] = userInfo
	response["contributions"] = graphs

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
