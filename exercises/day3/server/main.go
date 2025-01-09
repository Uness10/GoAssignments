package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}
type User struct {
	Id      int
	Name    string
	Email   string
	Age     int
	Address Address
}

var Users = map[int]User{}

func addUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Users[user.Id] = user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(Users)
	//json.NewEncoder(w).Encode(user)
}

/*
	func handlerPOST(w http.ResponseWriter, r *http.Request) {
		req := make(map[string]interface{})
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Using POST method!\n")
		fmt.Fprintln(w, "Your data:", req)
	}
*/
func QueryParamsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == http.MethodGet {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Query Params are : ", r.Form)
	}
}

func PathParamsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Hello, user %s", id)
}

func handlerPUT(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func main() {
	router := httprouter.New()
	router.GET("/user/:id", PathParamsHandler)
	router.GET("/welcome", QueryParamsHandler)
	router.POST("/add", addUser)
	router.PUT("/user/:id", handlerPUT)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
