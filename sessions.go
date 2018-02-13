package main

import (
	"net/http"
	"github.com/google/uuid"
)

func alreadyLoggedIn(r *http.Request) bool {
	//is there a session cookie?
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	//is user logged in?
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	//is user logged in?
	if alreadyLoggedIn(r) {
		//get cookie; error already checked
		c, _ := r.Cookie("session")
		//get user; error already checked
		un := dbSessions[c.Value]
		u, _ := dbUsers[un]
		return u
	}
	//if not logged in
	var u user
	return u

}

func signupProcess(w http.ResponseWriter, r *http.Request) {
	//do you have cookie?
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	//process form submission
	un := r.FormValue("username")
	f := r.FormValue("firstname")
	l := r.FormValue("lastname")

	if _, ok := dbUsers[un]; ok {
		http.Error(w, "Username already taken", http.StatusForbidden)
		return
	}

	dbSessions[c.Value] = un
	u := user{un, f, l}
	dbUsers[un] = u

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return

}
