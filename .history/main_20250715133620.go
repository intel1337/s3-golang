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
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		for _, u := range users {
			if u.Username == creds.Username && u.Password == creds.Password {
				next(w, r)
				return
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", authMiddleware(getUsers))
	http.ListenAndServe(":8080", nil)
}
