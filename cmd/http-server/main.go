package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

var users = []User{
	{"Massimo", "Costa", "massimoc@example.com"},
	{"John", "Doe", "doej@example.com"},
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/users", usersV1)
	mux.HandleFunc("/v2/users", usersV2)
	mux.HandleFunc("/v3/users", usersV3)
	mux.HandleFunc("/v4/users", usersV4)

	http.ListenAndServe(":8080", mux)
}

func usersV1(w http.ResponseWriter, r *http.Request) {
	// Simplest
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println("something went wrong: ", err)
	}
}

func usersV2(w http.ResponseWriter, r *http.Request) {
	// Read from request's headers
	foo := r.Header.Get("X-Foo")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Custom-Data", fmt.Sprintf("hello %s", foo))

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println("something went wrong: ", err)
	}
}

func usersV3(w http.ResponseWriter, r *http.Request) {
	// super simple authentication
	foo := r.Header.Get("X-Foo")
	if foo != "massimoc" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("I don't know you"))
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Custom-Data", fmt.Sprintf("hello %s", foo))

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println("something went wrong: ", err)
	}
}

func usersV4(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
	// super simple authentication
	foo := r.Header.Get("X-Foo")
	if foo != "massimoc" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("I don't know you"))
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("X-Custom-Data", fmt.Sprintf("hello %s", foo))

	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println("something went wrong: ", err)
	}
}
