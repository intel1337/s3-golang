package main

import (
	ID	string	`json:"id"`
	Username string `json:"username"`
	Permissions bool `json:"permissions"`
	Password string	`json:"password"`
}

var newUser = []User{
	{ID:1, Username: "root", Permissions: true, Password: "root"}
}




