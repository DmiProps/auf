package handlers

import (
	"net/http"
	"text/template"
)

// Index is handler for index page
func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("www/index.html")
	t.Execute(w, nil)
}
