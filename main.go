package main

import (
	"net/http"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	mux := http.NewServeMux()

	// Servir les fichiers statiques
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Servir la page HTML
	mux.HandleFunc("/", serveHTML)

	http.ListenAndServe(":8080", mux)
}
