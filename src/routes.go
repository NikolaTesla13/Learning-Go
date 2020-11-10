package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/logrusorgru/aurora"

	_ "github.com/go-sql-driver/mysql"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	data, _ := ioutil.ReadFile("pages/index.html")
	fmt.Fprintf(writer, string(data))
	fmt.Println("📬", Bold(Cyan("GET")), "request: '/'")
}

func getAllUsers(writer http.ResponseWriter, request *http.Request) {
	results, err := db.Query("SELECT * from users")
	fmt.Println("fetched")
	if err != nil {
		log.Print(err.Error())
	}
	var fetched_users []User
	for results.Next() {
		var user User
		err := results.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Verified)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(user.Name)
		fetched_users = append(fetched_users, user)
	}
	json.NewEncoder(writer).Encode(fetched_users)
	fmt.Println("📬", Bold(Cyan("GET")), "request: '/api/users/'")
}

func getSpecificUsers(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	for _, user := range users {
		if user.Id == id {
			json.NewEncoder(writer).Encode(user)
		}
	}
	fmt.Println("📬", Bold(Cyan("GET")), "request: '/api/users/{id}'")
}

func createNewUser(writer http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	fmt.Fprintf(writer, "%+v", string(requestBody)) //TODO
	fmt.Println("📬", Bold(Cyan("POST")), "request: '/api/users/'")
}
