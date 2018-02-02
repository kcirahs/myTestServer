package main

import (
	"io"
	"net/http"
)

func hf(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Internet!")
}

func hh(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Hannah!")
}

func main() {
	http.HandleFunc("/", hf)
	http.HandleFunc("/hannah", hh)
	http.ListenAndServe(":8080", nil)
}
