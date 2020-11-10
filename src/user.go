package main

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Verified bool   `json:"verified"`
	// TODO
}

var users []User
