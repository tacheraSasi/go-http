package main

import (
	"fmt"
	"net/http"
	"sync"
)
type User struct{
	Name string `json:"name"`
}

var usersCache = make(map[int]User)
var cacheMutex sync.RWMutex

func main(){
	// mux := http.NewServeMux()

	// mux.HandleFunc("/" ,handleRoot)
	// mux.HandleFunc("POST /users" ,createUser)
	// mux.HandleFunc("GET /users/{id}", getUser)
	// mux.HandleFunc("DELETE /users/{id}", deleteUser)

	//starting the server
	fmt.Println("server is running on http://localhost:8000")
	// http.ListenAndServe(":8000",mux)
	http.ListenAndServe(":8000",http.FileServer(http.Dir("/var/www/html")))
}
