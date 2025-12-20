package auth

import "fmt"

type Credentials struct {
	username string
	password string
}

func LoginWithCredentials(username string, password string) {
	fmt.Printf("login user using %s %s\n", username, password)
}
