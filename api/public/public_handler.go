package public

import (
	"net/http"
)

// Handler para servir arquivos estáticos do diretório public
func PublicHandle() http.HandlerFunc {
	fs := http.FileServer(http.Dir("./public"))
	return func(w http.ResponseWriter, r *http.Request) {
		// Remove "/public/" do início da URL para mapear corretamente o arquivo
		http.StripPrefix("/doc/", fs).ServeHTTP(w, r)
	}
}