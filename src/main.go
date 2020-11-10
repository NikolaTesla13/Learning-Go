package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	. "github.com/logrusorgru/aurora"
)

var db *sql.DB
var err error

func init_server() {
	fmt.Println("🚀", Bold(Magenta("Starting")), "the server...")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/api/users/", getAllUsers)
	router.HandleFunc("/api/users/{id}", getSpecificUsers)
	router.HandleFunc("/api/user/", createNewUser).Methods("POST")

	fmt.Println("✅", Bold(Green("Listening")), "on localhost:10000")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:10000)/project")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	users = []User{
		User{Id: "123", Name: "User1", Email: "user.1@gmail.com", Password: "wejfhew", Verified: true},
		User{Id: "345", Name: "User2", Email: "user.2@gmail.com", Password: "wejfhew", Verified: false},
	}

	go func() {
		init_server()
	}()
	<-killSignal

	fmt.Println("🛑", Bold(Red("Closing")), "the server...")

}
