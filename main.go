package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var DB *sql.DB

func main() {
	connstring := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "localadmin", "godzirras")
	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	}

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

	r.ParseForm()
	rows, err := DB.Query("INSERT INTO godzillas(name, height) VALUES ('" + r.FormValue("name") + "', '" + r.FormValue("height") + "')")
	fmt.Println(rows)
	fmt.Println(err)
	// http.ServeFile(w, r, "templates/godzirras.html")
}
