package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"github.com/mavricknz/ldap"

	"github.com/gorilla/mux"
)

var DB *sql.DB

// sample struct for passing information to the html templates
type Stuff struct {
	Blah   string
	Blue   string
	Errors NameErr
}

// struct that contains all the errors for validations
type NameErr struct {
	TooShort string
	TooLong  string
}

func main() {
	server := os.Getenv("SERVER")
	database := os.Getenv("DATABASE")
	password := os.Getenv("PASSWORD")
	userID := os.Getenv("USERID")

	connstring := fmt.Sprintf("server=%s;database=%s;password=%s;user id=%s;", server, database, password, userID)
	var err error
	DB, err = sql.Open("mssql", connstring)
	if err != nil {
		log.Fatal("err", err)
	}
	// assures the database is working
	err = DB.Ping()
	if err != nil {
		fmt.Println("Ping: ", err)
	}

	// establishing new router w/routes and handlers
	fmt.Println(DB)
	r := mux.NewRouter()
	r.HandleFunc("/", TokyoHandler).
		Methods("GET")
	r.HandleFunc("/godzirras", GodzirrasHandler).
		Methods("POST")
	// handle css links in html
	r.HandleFunc("/css", css).
		Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", r))
}

// Handler function for the home page
func TokyoHandler(w http.ResponseWriter, r *http.Request) {
	getEnvironment()
	render(w, "templates/tokyo.html", Stuff{Blue: "True"})
}

// Handler for the page that loads after submitting
func GodzirrasHandler(w http.ResponseWriter, r *http.Request) {
	// parse the form values
	r.ParseForm()
	name := r.FormValue("name")
	height := r.FormValue("height")

	// check if the inputs are valid
	isValid, errList := nameValidator(name)
	fmt.Println(isValid)
	fmt.Println(errList)

	// if the inputs are valid insert into the database, otherwise re render the same page with the errors
	if isValid == true {
		DB.Query("INSERT INTO godzillas(name, height) VALUES ('" + name + "', '" + height + "')")
	} else {
		render(w, "templates/tokyo.html", Stuff{Blue: "True", Errors: errList})
	}
}

// serving css pages
func css(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/css/styles.css")
}

func ErrorChecker() string {
	// using golang time package to create a dynamic seed for randomness
	now := time.Now()
	nanos := int64(now.Nanosecond())
	rand.Seed(nanos)

	// bunch of different vague responses we will be pulling from
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
	// use the random integer function to pull a random index from answers.
	return "Magic 8-Ball says: " + answers[rand.Intn(len(answers))]
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	// creating a funcMap that we can use to add custom functions to the template
	funcMap := template.FuncMap{
		"ErrorChecker": ErrorChecker,
		"rando":        rando,
		"sayMuch":      sayMuch,
		"epic":         epicImages,
	}

	// creating a new template and including the functions that we added to the func map
	tmpl, err := template.New("tokyo.html").Funcs(funcMap).ParseFiles(filename)
	if err != nil {
		log.Fatal("STuff and such:", err)
	}
	// render the template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal("MOAR STUFF: ", err)
	}
}

// random number generator
func rando() int {
	now := time.Now()
	nanos := int64(now.Nanosecond())
	rand.Seed(nanos)
	return rand.Intn(10)
}

// returns a string
func sayMuch(repeat int) string {
	return "I say a lot " + strconv.Itoa(repeat) + " times"
}

// randomly assigns an image url
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

//
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

func getEnvironment() {
	// first := os.Environ()
	// for key, val := range first {
	// 	if val == "josh=bad hair" {
	// 		fmt.Println(key)
	// 	}
	// }

	first := os.Getenv("josh")
	second := os.Getenv("foo")
	third := os.Getenv("stinnette")

	fmt.Println("Josh = ", first)
	fmt.Println("foo = ", second)
	fmt.Println("stinnette = ", third)
}

func Validator() string {
	return "ERR DOOD"
}
