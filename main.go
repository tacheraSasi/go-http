package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
var PORT = ":8000"
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("GET /",showInfo)
	mux.HandleFunc("GET /sites",serveFile)
	fmt.Println("Server is running smoothly at http://localhost:"+PORT)
	err := http.ListenAndServe(PORT,mux)
	if err != nil{
		log.Fatal("Something went wrong while running the server",err)
	}
}

func showInfo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Current Time:",time.Now())
}

func serveFile(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"frontend/index.html")
}