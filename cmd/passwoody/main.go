package main

import (
	"fmt"

	"github.com/crafty-codes/passwoody/internal/hashit"
)

func main() {
	fmt.Print("Please enter your Masterpassword: ")
	var masterpassword string
	fmt.Scanln(&masterpassword)

	fmt.Print("Please enter website: ")
	var password string
	fmt.Scanln(&password)

	hashit.Hasher(masterpassword + password)
}
