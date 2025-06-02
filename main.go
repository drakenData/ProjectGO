package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	// Rota do cabeçalho para a home page ( root URL )
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // Arquivos Statico do CSS //
	http.HandleFunc("/", home)
	http.HandleFunc("/projects", projects)
	http.HandleFunc("/skills", skills)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Está se iniciando no http://localhost:8080")

	// Iniciando o Servidor na Porta 8080 e escutando as requisições
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// home function cabeçalho para pedir requisções para home page
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")

}
func projects(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "projects.html")

}
func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html")

}
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")

}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")

}
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parsing espefica uma template que passa o imput (home< projeto < etc.)
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
