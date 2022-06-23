package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

func main() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/v4/users", bytes.NewBufferString(""))
	req.Header.Add("X-Foo", "massimoc")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error loading data: %v\n", err)
		os.Exit(1)
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		body, _ := io.ReadAll(r.Body)
		fmt.Printf("unexpected status (%s): %s\n", r.Status, string(body))
		os.Exit(1)
	}

	var users []User

	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		fmt.Printf("error decoding JSON response: %s\n", err)
	}
	fmt.Printf("Users: %+v\n", users)
}
