package main

import (
	"fmt"

	"github.com/nirabyte/basics/auth"
	"github.com/nirabyte/basics/user"
)

func main() {
	auth.LoginWithCredentials("nirabyte", "secret ")
	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "User",
	}
	fmt.Println(user.Email, user.Name)
}
