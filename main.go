package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", TokyoHandler).
		Methods("GET")
	r.HandleFunc("/godzirras", GodzirrasHandler).
		Methods("POST")
	log.Fatal(http.ListenAndServe(":9001", r))
}

func TokyoHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/tokyo.html")
}

func GodzirrasHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/godzirras.html")
}
