package main

import (
	"fmt"
	"log"
	"net/http"

	"api_git_leet_duo/api/duo"
	"api_git_leet_duo/api/git/handler"
	"api_git_leet_duo/api/leet"
	"api_git_leet_duo/api/public"
)

func main() {
	// Serve static files from public directory
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// API routes - matching vercel.json configuration
	http.HandleFunc("/api/git/user", handler.GitUser)
	http.HandleFunc("/api/git/repos", handler.GitRepos)
	http.HandleFunc("/api/git/langs", handler.GitLangs)
	http.HandleFunc("/api/git/streak", handler.GitStreak)
	http.HandleFunc("/api/git/commit", handler.GitCommit)

	// Duolingo API
	http.HandleFunc("/api/duo/user", duo.DuoUser)

	// LeetCode API
	http.HandleFunc("/api/leet/user", leet.LeetUser)

	// Documentation route
	http.HandleFunc("/api/doc/", public.PublicHandle)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("API Documentation: http://localhost:8080/api/doc/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
