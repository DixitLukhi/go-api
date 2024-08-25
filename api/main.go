package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

type errorResponse struct {
	StatusCode int
	Message    string
}

var users = map[string]User{}

func main() {
	http.HandleFunc("/createuser", addUser)
	http.HandleFunc("/users", getUsers)

	fmt.Println("Server started")
	log.Fatalf("Server crash : %v\n", http.ListenAndServe(":4000", nil))

}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errr := errorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Wrong payload",
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	users[user.Name] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	fmt.Println("Users are : ", user)
	return
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errr := errorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Something went wrong",
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
	return
}
