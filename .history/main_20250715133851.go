package main

import (
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

// Simple authentication: checks for Basic Auth with root/root
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		for _, user := range users {
			if user.Username == username && user.Password == password && user.Permissions {
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
	http.HandleFunc("/addfile", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi admin"))
	}))
	http.ListenAndServe(":8080", nil)
}
