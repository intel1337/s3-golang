package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Permissions bool  `json:"permissions"`
	Password   string `json:"password"`
}

var users = []User{
	{ID: 1, Username: "root", Permissions: true, Password: "root"},
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		for _, u := range users {
			if u.Username == username && u.Password == password {
				next(w, r)
				return
			}
		}
		http.Error(w,
