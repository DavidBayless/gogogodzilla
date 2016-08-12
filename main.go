package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type Stuff struct {
	Blah string
}

func main() {
	connstring := fmt.Sprintf("server=%s;database=%s;password=%s;user id=%s;", "a9sst001.allstate.com", "A9D_Alfred_EM_DEV", "A9Adm00DEV", "A9U_Admin_DEV")
	var err error
	DB, err = sql.Open("mssql", connstring)
	if err != nil {
		log.Fatal("err", err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Ping: ", err)
	}
	fmt.Println(DB)
	// funcMap := template.FuncMap {
	// 	"Validator": Validator,
	// }
	// template.Funcs(funcMap)
	r := mux.NewRouter()
	r.HandleFunc("/", TokyoHandler).
		Methods("GET")
	r.HandleFunc("/godzirras", GodzirrasHandler).
		Methods("POST")
	log.Fatal(http.ListenAndServe(":9001", r))
}

func TokyoHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/tokyo.html", Stuff{Blah: "Blue"})
}

func GodzirrasHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	DB.Query("INSERT INTO godzillas(name, height) VALUES ('" + r.FormValue("name") + "', '" + r.FormValue("height") + "')")
	//fmt.Println(rows)
	//fmt.Println(err)
	// http.ServeFile(w, r, "templates/godzirras.html")
}

func Validator() string {
	return "ERR DOOD"
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Blahhh")
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
