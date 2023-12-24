package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "POST request successful\n")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World from Go!!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on port :8089 ...")

	if err := http.ListenAndServe(":8089", nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	templatePath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
	}
}
