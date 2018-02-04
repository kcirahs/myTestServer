package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

var emails []string

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.html", emails)
}

func main() {
	emails = append(emails, "someguy@gmail.com")

	http.Handle("/", http.HandlerFunc(rootHandler))
	http.Handle("/about/", http.HandlerFunc(aboutHandler))
	http.Handle("/contact/", http.HandlerFunc(contactHandler))

	http.ListenAndServe(":8080", nil)
}
