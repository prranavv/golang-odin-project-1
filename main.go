package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/about", about)
	r.Get("/contact", contact)
	r.Get("/", mainpage)
	r.Get("/{fileName}", serveHTML)
	http.ListenAndServe(":3000", r)
}

func mainpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")

	// Check if the requested file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Serve the custom 404 page if the file doesn't exist
		http.ServeFile(w, r, "static/error.html")
		return
	}

	// Serve the requested HTML file
	http.ServeFile(w, r, fileName)
}
