package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/DmiProps/auf/database"
	"github.com/DmiProps/auf/handlers"
	"github.com/DmiProps/auf/settings"
)

func main() {

	settings.ReadSettings()

	database.GetConnect()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/signup", handlers.Signup)

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
