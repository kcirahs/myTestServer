package main

import (
	"html/template"
	"net/http"
	"log"
	"math/rand"
)

var tmpl *template.Template

type addition struct {
	N1  int
	N2  int
	Sum int
	AnswerField string
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func generateAdd() (addition) {
	n1 := rand.Intn(10)
	n2 := rand.Intn(10)
	sum := n1 +n2
	return addition{n1, n2, sum, ""}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	a := generateAdd()
	a.AnswerField = r.FormValue("answer")

	tmpl.ExecuteTemplate(w, "contact.html", a)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/contact/", contactHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
