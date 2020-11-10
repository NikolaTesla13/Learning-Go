package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	. "github.com/logrusorgru/aurora"
)

func init_server() {
	fmt.Println("🚀", Bold(Magenta("Starting")), "the server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	fmt.Println("✅", Bold(Green("Listening")), "on localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	go func() {
		init_server()
	}()
	<-killSignal

	fmt.Println("🛑", Bold(Red("Closing")), "the server...")
}
