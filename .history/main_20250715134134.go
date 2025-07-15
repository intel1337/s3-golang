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

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticates using JSON credentials in the request body
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		for _, user := range users {
			if user.Username == creds.Username && user.Password == creds.Password && user.Permissions {
				next(w, r)
				return
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	http.HandleFunc("/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	http.HandleFunc("/add-file", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi admin"))
	}))
	http.ListenAndServe(":8080", nil)
}
