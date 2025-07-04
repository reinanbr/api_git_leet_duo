package main

import (
	"fmt"
	"log"
	"net/http"

	"api_git_leet_duo/api/public"
)

func main() {

	http.HandleFunc("/api/doc/", public.PublicHandle().ServeHTTP)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
