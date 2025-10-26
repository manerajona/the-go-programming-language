package main

import (
	"fmt"
)

func main() {
	var firstname string
	var lastname string
	fmt.Print("Enter your firstname and lastname: ")

	inputs, _ := fmt.Scanf("%s %s", &firstname, &lastname)
	switch inputs {
	case 0:
		fmt.Println("Missing firstname and lastname")
	case 1:
		fmt.Printf("Missing lastname")
	case 2:
		fmt.Printf("Hello %s %s\n", lastname, firstname)
	}
}
