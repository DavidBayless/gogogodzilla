package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/mavricknz/ldap"

	"github.com/gorilla/mux"
)

var DB *sql.DB

type Stuff struct {
	Blue   string
	Errors NameErr
}

type NameErr struct {
	TooShort string
	TooLong  string
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
	r.HandleFunc("/css", css).
		Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", r))
}

func TokyoHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/tokyo.html", Stuff{Blue: "True"})
}

func GodzirrasHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	name := r.FormValue("name")
	height := r.FormValue("height")
	isValid, errList := nameValidator(name)
	fmt.Println(isValid)
	fmt.Println(errList)

	if isValid == true {
		DB.Query("INSERT INTO godzillas(name, height) VALUES ('" + name + "', '" + height + "')")
	} else {
		render(w, "templates/tokyo.html", Stuff{Blue: "True", Errors: errList})
	}
}

func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/css/styles.css")
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
	return "Magic 8-Ball says: " + answers[rand.Intn(len(answers))]
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	funcMap := template.FuncMap{
		"ErrorChecker": ErrorChecker,
		"rando":        rando,
		"sayMuch":      sayMuch,
		"epic":         epicImages,
	}

	tmpl, err := template.New("tokyo.html").Funcs(funcMap).ParseFiles(filename)
	if err != nil {
		log.Fatal("STuff and such:", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal("MOAR STUFF: ", err)
	}
}

func rando() int {
	now := time.Now()
	nanos := int64(now.Nanosecond())
	rand.Seed(nanos)
	return rand.Intn(10)
}

func sayMuch(repeat int) string {
	return "I say a lot " + strconv.Itoa(repeat) + " times"
}

func epicImages() string {
	now := time.Now()
	nanos := int64(now.Nanosecond()) // Try changing this number!
	rand.Seed(nanos)
	images := []string{
		"http://i0.kym-cdn.com/photos/images/original/000/452/840/d73.jpg",
		"http://orig05.deviantart.net/7153/f/2011/053/1/7/teddy_roosevelt_vs__bigfoot_by_sharpwriter-d3a72w4.jpg",
		"http://36.media.tumblr.com/ccbebafa27c043438adb15cbc0615bac/tumblr_nqy4t2Ln0i1qfaphzo1_1280.jpg",
	}
	return images[rand.Intn(len(images))]
}

func nameValidator(name string) (bool, NameErr) {
	errorList := NameErr{TooLong: "", TooShort: ""}
	isValid := true
	if len(name) < 3 {
		isValid = false
		errorList.TooShort = "Your name is too short!"
	}

	if len(name) > 20 {
		isValid = false
		errorList.TooLong = "Your name is too Long!"
	}
	return isValid, errorList
}

//LDAP FUNC EXAMPLE ----
func auth(username string, password string) {
	//server string is where you put the server string without the 'LDAP://'
	//000 is where your port goes
	connection := ldap.NewLDAPConnection("server string", uint16(000))
	err := connection.Connect()
	fmt.Println("errrrrrrrr: ", err)
	defer connection.Close()
	err = connection.Bind(username, password)
	if err != nil {
		fmt.Println("You suck, ", err)
	}
}
