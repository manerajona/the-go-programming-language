package main

import "fmt"

var s string

func main() {

	s = "Go is all about types"

	fmt.Printf("s value is %q\n", s)
	fmt.Printf("s type is %T\n", s)

	// String as a slice
	fmt.Printf("s lenght is %d\n", len(s))
	fmt.Println(s[0:2], "is cool")
	fmt.Println(s[:2], "is cool")
	fmt.Println("Javascript is NOT", s[10:])

}
