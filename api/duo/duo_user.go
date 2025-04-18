package duo

import (
	"api_git_leet_duo/api/duo/tools"
	"encoding/json"
	"net/http"

)


type UserData struct {
	Fullname     string       `json:"fullname"`
	Username     string       `json:"username"`
	Avatar       string       `json:"avatar"`
	CreationDate string       `json:"creation_date"`
	Courses      []struct {
		Title string `json:"title"`
		XP    int    `json:"xp"`
	} `json:"courses"`
	XPByLanguage []tools.LanguageXP `json:"xp_by_language"` // <-- Novo campo
}


type APIResponse struct {
	Users []UserData `json:"users"`
}



func DuoUser(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "Parâmetro 'user' é obrigatório", http.StatusBadRequest)
		return
	}

	userData, err := tools.FetchDuolingoUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userData)
}


