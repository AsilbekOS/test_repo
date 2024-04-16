package main

import (
	"fmt"

	"github.com/AsilbekOS/test_repo/user"
)

func main() {
	var name string
	var age int
	var email string

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Age: ")
	fmt.Scan(&age)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	errors := user.ValidateUser(name, age, email)

	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	}
}
