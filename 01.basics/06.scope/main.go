package main

import "fmt"

func main() {
	// Global scope
	fmt.Println("global:", global)
	fmt.Println("Outside:", Outside)

	// Function scope (local)
	internal := 10
	fmt.Println("internal:", internal)

	// Closures; Block scope (local)
	{
		// variable shadowing
		internal := 0
		fmt.Println("internal:", internal)
		internal++
		fmt.Println("internal:", internal)
		internal++
		fmt.Println("internal:", internal)
		{
			// variable shadowing again
			internal := 10
			fmt.Println("internal:", internal)
			internal++
			fmt.Println("internal:", internal)
			internal++
			fmt.Println("internal:", internal)
		}
		internal++
		fmt.Println("internal:", internal)
	}
	fmt.Println("internal:", internal)
}

// global identifier has PRIVATE access since starts with lowercase
var global int

/* PUBLIC access constants:
When importing a package you can access only its exported identifiers.
An identifier is exported if it begins with a capital letter.
*/

// Outside represents a string constant "Hello"
const Outside = "Hello"
