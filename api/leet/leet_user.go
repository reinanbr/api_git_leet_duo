package leet

import (
	"api_git_leet_duo/api/leet/tools"
	"encoding/json"


	"net/http"
)



func LeetUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	userData, err := tools.GetUserData(username)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}

