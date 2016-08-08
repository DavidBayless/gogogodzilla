package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"math/rand"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type Stuff struct {
	Blue string
}

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
	// http.ServeFile(w, r, "templates/tokyo.html")
	render(w, "templates/tokyo.html", Stuff{Blue: "True"})
}

func GodzirrasHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	DB.Query("INSERT INTO godzillas(name, height) VALUES ('" + r.FormValue("name") + "', '" + r.FormValue("height") + "')")
	//fmt.Println(rows)
	//fmt.Println(err)
	// http.ServeFile(w, r, "templates/godzirras.html")
}

func ErrorChecker() string {
	now := time.Now()
	nanos := int64(now.Nanosecond()) // Try changing this number!
	rand.Seed(nanos)
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	return "Magic 8-Ball says: "+ answers[rand.Intn(len(answers))]
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	funcMap := template.FuncMap{
		"ErrorChecker": ErrorChecker,
	}

	tmpl, err := template.New("tokyo.html").Funcs(funcMap).ParseFiles(filename)
	if err != nil {
		log.Fatal("STuff and such:", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("MOAR STUFF: ", err)
	}
}
