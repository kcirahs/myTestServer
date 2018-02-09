package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/kcirahs/myTestServer/utils"
	"fmt"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func flashCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		m := r.Form
		PreExample := utils.ParsePrevious(m["PreExample"][0])
		fmt.Println(PreExample)
		tmpl.ExecuteTemplate(w, "answer.html", PreExample)
		return
	}
	Example := utils.GenerateDiff()
	//Example := utils.GenerateAdd()
	tmpl.ExecuteTemplate(w, "flashCard.html", Example)
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/flashCard/", flashCardHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
