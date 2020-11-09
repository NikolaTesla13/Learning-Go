package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("pages/index.html")
	fmt.Fprintf(w, string(data))
	fmt.Println("GET request: '/'")
}
