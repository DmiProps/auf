package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DmiProps/auf/database"
	"github.com/DmiProps/auf/handlers"
	"github.com/DmiProps/auf/settings"
)

func main() {

	settings.ReadSettings()

	database.Connect()
	defer settings.DbConnect.Close(context.Background())

	addHTTPRouter()

	http.ListenAndServe(":"+settings.AppSettings.Port, nil)

}

func addHTTPRouter() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/signup", handlers.Signup)
	r.HandleFunc("/activation", handlers.ActivateViaEmail)

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

}
