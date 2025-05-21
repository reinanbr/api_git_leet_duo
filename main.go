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

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
