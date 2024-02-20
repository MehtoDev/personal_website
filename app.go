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
var resources embed.FS
var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", ContentHandler("page_index"))
	http.HandleFunc("/public/css/styles", FileHandler("/public/css/styles.css"))
	http.HandleFunc("/public/img/Github-Logo", FileHandler("/public/img/Github-Logo.png"))
	http.HandleFunc("/public/img/Github", FileHandler("/public/img/Github.png"))
	http.HandleFunc("/public/img/Linkedin-Logo", FileHandler("/public/img/Linkedin-Logo.png"))
	http.HandleFunc("/public/scripts/class-tools", FileHandler("/public/scripts/class-tools.js"))
	http.HandleFunc("/public/scripts/htmx", FileHandler("/public/scripts/htmx.js"))

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

func FileHandler(name string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, fmt.Sprintf("%s", name))
	}
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
