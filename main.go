package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/kcirahs/myTestServer/utils/flashCard"
)

type user struct {
	Username string
	First    string
	Last     string
}

var tmpl *template.Template
var dbUsers = map[string]user{}      //userID, user
var dbSessions = map[string]string{} //sessionID, userID

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/logon/", signup)
	http.HandleFunc("/signup/", signup)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/flashCard/", flashCardHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//process signup data
	if r.Method == http.MethodPost {
		signupProcess(w, r)
		return
	}
	tmpl.ExecuteTemplate(w, "signup.html", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tmpl.ExecuteTemplate(w, "index.html", u)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func flashCardHandler(w http.ResponseWriter, r *http.Request) {
	//if showing answer
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		m := r.Form
		PreExample := flashCard.ParsePrevious(m["PreExample"][0])
		tmpl.ExecuteTemplate(w, "answer.html", PreExample)
		return
	}
	//creating new card
	var Example flashCard.MathProblem
	switch r.URL.Path {
	case "/flashCard/diff":
		Example = flashCard.GenerateDiff()
	case "/flashCard/add":
		Example = flashCard.GenerateAdd()
	default:
		Example = flashCard.GenerateAdd()
	}
	tmpl.ExecuteTemplate(w, "flashCard.html", Example)
}
