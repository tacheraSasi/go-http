package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type FormData struct {
	Name  string
	Email string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse and execute the template
	tmplPath := filepath.Join("./templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}
	t.Execute(w, data)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Render the form
	renderTemplate(w, "index.html", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Handle POST requests only
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		log.Println("Form parse error:", err)
		return
	}

	// Retrieve form fields
	data := FormData{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	// Render the thank-you template with form data
	renderTemplate(w, "thankyou.html", data)
}

func main() {
	// Set up routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", formHandler)
	mux.HandleFunc("/submit", submitHandler)

	// Start the server
	log.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
