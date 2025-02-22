package main

import "net/http"

func main(){
	r := http.NewServeMux()
	r.HandleFunc("/healthcheck", healthCheck)


}

func healthCheck(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}