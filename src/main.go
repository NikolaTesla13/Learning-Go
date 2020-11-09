package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

func init_server() {
	fmt.Println("🚀 Starting the server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	fmt.Println("✅ Listening on localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	go func() {
		init_server()
	}()
	<-killSignal

	fmt.Println("🛑 Closing the server...")
}
