package main

import "net/http"

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
		// Example: allow all requests for now
		next(w, r)
	}
}
func main() {
	// Example: Start an HTTP server (add your handlers here)
	http.HandleFunc("/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	http.ListenAndServe(":8080", nil)