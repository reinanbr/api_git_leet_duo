package main

import (
	"fmt"
	"log"
	"net/http"

	"api_git_leet_duo/api/git/handler"
)

func main() {
	http.HandleFunc("/git/user", handler.GitUser)
	http.HandleFunc("/git/repos", handler.GitRepos)
	http.HandleFunc("/git/langs", handler.GitLangs)
	http.HandleFunc("/git/streak", handler.GitStreak)
	http.HandleFunc("/git/commit", handler.GitCommit)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
