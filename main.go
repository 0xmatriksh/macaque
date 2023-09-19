package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(user)
	}
	fmt.Printf("Hello, %s. Welcome", user.Username)
}
