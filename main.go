package main

import (
	"html/template"
	"net/http"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var tmpl *template.Template

type addition struct {
	N1  int
	N2  int
	Sum int
	//AnswerField string
}

var AddExample addition
var r1 *rand.Rand

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	r1 = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func generateAdd() (addition) {
	n1 := r1.Intn(10)
	n2 := r1.Intn(10)
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
