package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed templates/*
var templates embed.FS
var t = template.Must(template.ParseFS(templates, "templates/*"))

//go:embed all:assets
var assets embed.FS

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", ContentHandler("page_index"))
	http.Handle("/assets/", http.FileServer(http.FS(assets)))

	//Full page
	http.HandleFunc("/home", ContentHandler("page_index"))
	http.HandleFunc("/projects", ContentHandler("page_projects"))
	http.HandleFunc("/projects/gamedev", ContentHandler("page_gamedev"))

	//HTMX handlers
	http.HandleFunc("/nav/top_level", NavHandler(1))
	http.HandleFunc("/nav/projects", NavHandler(2))
	http.HandleFunc("/nav/gamedev", NavHandler(3))
	http.HandleFunc("/content/footer", ContentHandler("footer"))
	http.HandleFunc("/content/home", ContentHandler("content_home"))
	http.HandleFunc("/content/projects", ContentHandler("content_projects"))
	http.HandleFunc("/content/gamedev", ContentHandler("content_gamedev"))
	http.HandleFunc("/content/about", ContentHandler("content_about"))
	http.HandleFunc("/content/chess", ContentHandler("content_chess"))
	http.HandleFunc("/content/godot_platformer", ContentHandler("content_godot_platformer"))
	http.HandleFunc("/content/godot_fps_controller", ContentHandler("content_godot_fps_controller"))

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func NavHandler(i int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]int{
			"block": i,
		}
		t.ExecuteTemplate(w, "nav.html.tmpl", data)
	}
}

func ContentHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, fmt.Sprintf("%s.html.tmpl", name), nil)
	}
}
