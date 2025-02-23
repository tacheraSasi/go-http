package main

import (
	"fmt"
	"net/http"
)

var print = fmt.Println

func main(){
	r := http.NewServeMux()
	r.HandleFunc("/healthcheck", healthCheck)
	r.HandleFunc("/api/register",SignUp)

	print("Server is running on http://localhost:8080")
	
	err := http.ListenAndServe(":8080", r)
	if err != nil {	
		panic(err)
	}
}

func healthCheck(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("OK\n"))
	w.WriteHeader(http.StatusOK)
}