package handler

import (
	"api_git_leet_duo/api/git/tools/languages"
	"encoding/json"
	"fmt"
	"net/http"
)

func GitLangs(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	langs, err := languages.FetchUserLangsFull(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving graphsLangs: %v", err), http.StatusInternalServerError)
		return
	}

	// Monta a resposta JSON
	response := make(map[string]interface{})
	langs_percent,totalSize := languages.CalculateLanguagePercentage(langs)
	response["total_size"] = totalSize
	response["percent_lang"] = langs_percent

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
