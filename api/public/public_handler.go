package public

import (
	"net/http"
	"os"
)

// Handler para servir arquivos estáticos do diretório public
func PublicHandle(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./public/index.html")
	if err != nil {
		http.Error(w, "Arquivo não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}