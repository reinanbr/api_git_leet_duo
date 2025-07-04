package public

import (
	"net/http"
	"fmt"
)

// Handler para servir arquivos estáticos do diretório public
func PublicHandle(w http.ResponseWriter, r *http.Request) {
	// Verifica se o arquivo requisitado existe
	if _, err := http.Dir("./public").Open(r.URL.Path[len("/api/doc/"):]); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Arquivo não encontrado: %v | path: %v"}`, err, r.URL.Path), http.StatusNotFound)
		fmt.Println("Erro ao servir arquivo:", err)
		return
	}

	fs := http.FileServer(http.Dir("./public"))
	http.StripPrefix("/api/doc/", fs).ServeHTTP(w, r)
	fmt.Println("Request received for:", r.URL.Path)
}