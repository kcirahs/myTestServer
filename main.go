package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

var emails []string

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.gohtml", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "contact.gohtml", emails)
}

func main() {
	emails = append(emails, "someguy@gmail.com")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/contact/", contactHandler)

	http.ListenAndServe(":8080", nil)
}
