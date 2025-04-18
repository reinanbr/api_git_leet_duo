package main

import (
	"fmt"
	"net/http"
//	git "api_git_leet_duo/api/git"
	duo "api_git_leet_duo/api/duo"
//	leet "api_git_leet_duo/api/leet"
)


func main() {
	// Define a rota para a API
	http.HandleFunc("/github", duo.DuoUser)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}