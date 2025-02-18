package handler

import (
	"api_git_leet_duo/api/git/tools/user"
	"encoding/json"
	"fmt"
	"net/http"
)

func GitUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("user")
	if username == "" {
		http.Error(w, "Missing 'user' parameter", http.StatusBadRequest)
		return
	}

	userInfo, errUser := user.FetchUserData(username)
	if errUser != nil {
		http.Error(w, fmt.Sprintf("Error retrieving graphsUser: %v", errUser), http.StatusInternalServerError)
		return
	}

	// Monta a resposta JSON
	response := make(map[string]interface{})
	response["user"] = userInfo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
