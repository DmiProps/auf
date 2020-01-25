package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("www/index.html")
		t.Execute(w, nil)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	r.PathPrefix("/www").Handler(
		http.StripPrefix(
			"/www",
			http.FileServer(http.Dir("./www"))))
	r.PathPrefix("/css").Handler(
		http.StripPrefix(
			"/css",
			http.FileServer(http.Dir("./www/css"))))
	r.PathPrefix("/images").Handler(
		http.StripPrefix(
			"/images",
			http.FileServer(http.Dir("./www/images"))))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
}
