package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Permissions bool   `json:"permissions"`
	Password    string `json:"password"`
}

var users = []User{
	{ID: 1, Username: "root", Permissions: true, Password: "root"},
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"Username"`
			Password string `json:"Password"`
		}
		if
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", authMiddleware(getUsers))
	http.ListenAndServe(":8080", nil)
}
