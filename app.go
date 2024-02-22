package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed all:templates
var templates embed.FS
var t = template.Must(template.ParseFS(templates, "templates/*.gohtml", "templates/*/*.gohtml"))

//go:embed all:assets
var assets embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/assets/", http.FileServer(http.FS(assets)))

	//HTMX handlers
	http.HandleFunc("/nav/", NavHandler)
	http.HandleFunc("/content/", ContentHandler)

	http.HandleFunc("/", PageHandler)
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func NavHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"block": r.URL.Path,
	}
	_ = data
	t.ExecuteTemplate(w, "nav", data)
}

func ContentHandler(w http.ResponseWriter, r *http.Request) {
	destination := strings.Replace(r.URL.Path, "/content/", "", -1)
	t.ExecuteTemplate(w, fmt.Sprintf("article/%s", destination), nil)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	destination := r.URL.Path
	if last := len(destination) - 1; last >= 1 && destination[last] == '/' {
		destination = destination[:last]
	}
	data := map[string]string{
		"block": destination,
	}
	t.ExecuteTemplate(w, "page", data)
}
