package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template
var tmplCSS *template.Template

var emails []string

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	tmplCSS = template.Must(template.ParseGlob("templates/*.css"))
}

func styleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css; charset=utf-8")
	tmplCSS.Execute(w, nil)
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

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/main.css", styleHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/contact/", contactHandler)

	http.ListenAndServe(":8080", nil)
}
