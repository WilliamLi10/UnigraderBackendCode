package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
    "fmt"
)
func main() {
    r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r, "public/index.html")
	})

    r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
        email := r.FormValue("email")
        password := r.FormValue("password")
    })

    http.ListenAndServe(":8080", r)
}
