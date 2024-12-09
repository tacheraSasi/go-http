package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)


func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"endpoint reached succesfully")

}
//handler to get a user
func getUser(w http.ResponseWriter, r *http.Request){
	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil{
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	cacheMutex.RLock()
	user,ok := usersCache[id]
	cacheMutex.RUnlock() 
	response := "User with id: "+ r.PathValue("id") +" Was not found"

	if !ok {
		http.Error(
			w,
			response,
			http.StatusNotFound,
		)
		return
	}

	

	w.Header().Set("Content-Type","application/json")

	retrievedUser,err := json.Marshal(user)
	if err != nil{
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(retrievedUser)
}

//handlerr to create users
func createUser(w http.ResponseWriter,r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if user.Name == ""{
		http.Error(
			w,
			"Username is empty",
			http.StatusBadRequest,
		)
	}

	cacheMutex.Lock()
	usersCache[len(usersCache)+1] = user
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

//handler to delete the user
func deleteUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil{
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
	}
	response := "User with id: "+ r.PathValue("id") +" Was not found"

	//checking if the user exists
	if _, ok := usersCache[id]; !ok {
		http.Error(
			w,
			response,
			http.StatusNotFound,
		)
		return
		
	}

	//deleting the user
	cacheMutex.Lock()
	delete(usersCache,id)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
