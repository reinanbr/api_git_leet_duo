package public

import (
	"net/http"
)

// Handler para servir arquivos estáticos do diretório public
func PublicHandle(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./public"))
	// Remove "/public/" do início da URL para mapear corretamente o arquivo
	http.StripPrefix("/api/doc/", fs).ServeHTTP(w, r)
}