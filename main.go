package main

import (
	"html/template"
	"net/http"
	"log"
	"math/rand"
	"strconv"
)

var tmpl *template.Template

type addition struct {
	N1  int
	N2  int
	Sum int
	//AnswerField string
}

var AddExample addition

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func generateAdd() (addition) {
	n1 := rand.Intn(10)
	n2 := rand.Intn(10)
	sum := n1 + n2
	return addition{n1, n2, sum}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func flashCardHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		sum, _:= strconv.Atoi(r.FormValue("sum"))
		n1, _:= strconv.Atoi(r.FormValue("n1"))
		n2, _:= strconv.Atoi(r.FormValue("n2"))
		AddExample = addition{n1, n2, sum}
		tmpl.ExecuteTemplate(w, "answer.html", AddExample)
		return
	}
	//AddExample.AnswerField = r.FormValue("answer")
	AddExample = generateAdd()
	tmpl.ExecuteTemplate(w, "flashCard.html", AddExample)
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/flashCard/", flashCardHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
