package main

import (
	"html/template"
	"net/http"

	"github.com/olivere/vite"
)

var tmpl *template.Template
var v vite.Generator

func main() {
	// Any template engine can be used
	tmpl = template.Must(template.ParseFiles("resources/index.tmpl"))

	v = vite.Generator{
		ViteURL:    "http://localhost:5173",
		Entrypoint: "resources/main.js",
	}

	// Use any web framework
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	pageData := struct {
		Vite vite.Generator
	}{
		Vite: v,
	}

	tmpl.Execute(response, pageData)
}
