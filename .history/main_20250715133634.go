package main

import (
	"encoding/json"
	"net/http"
	Permissions bool `json:"permissions"`
	Password string	`json:"password"`
}

var newUser = []User{
	{ID:1, Username: "root", Permissions: true, Password: "root"}
}




