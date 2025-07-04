package public

import (
	"net/http"
	"os"
)

// Handler para servir arquivos estáticos do diretório public
func PublicHandle(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./public/index.html")
	if err != nil {
		dirEntries, dirErr := os.ReadDir(".")
		if dirErr != nil {
			http.Error(w, "Arquivo não encontrado e não foi possível listar o diretório", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Arquivo não encontrado. Arquivos disponíveis:\n"))
		for _, entry := range dirEntries {
			w.Write([]byte(entry.Name() + "\n"))
		}
		http.Error(w, "Arquivo não encontrado", http.StatusNotFound)
		
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}