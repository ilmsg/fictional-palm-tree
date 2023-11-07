package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port = 7001

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").ParseFiles(
			"./templates/layout.html",
			"./templates/index.html",
		))
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	r.Get("/page", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").ParseFiles(
			"./templates/layout.html",
			"./templates/page.html",
		))
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	r.Get("/post", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("").ParseFiles(
			"./templates/layout.html",
			"./templates/post.html",
		))
		tmpl.ExecuteTemplate(w, "layout", nil)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		log.Fatalf("error on listen: *:%d", port)
	}
}
