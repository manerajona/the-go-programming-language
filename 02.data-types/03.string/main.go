package main

import "fmt"

var s string

func main() {

	fmt.Printf("default value of string is %q\n", s)

	s = "Go is all about types"

	fmt.Printf("s value is %q\n", s)
	fmt.Printf("s type is %T\n", s)

	// String as a slice
	fmt.Printf("s lenght is %d\n", len(s))
	fmt.Println(s[0:2], "is cool")
	fmt.Println(s[:2], "is cool")
	fmt.Println("Javascript is NOT", s[10:])

	// Back quotes
	s = `"Go is cool"`
	fmt.Printf("%s\n", s)

	s = `*******
*Go   *
*is   *
*cool *
*******`
	fmt.Printf("%s\n", s)

	// Single quotes (for byte or rune)
	// For try more symbols: https://www.w3schools.com/charsets/ref_utf_symbols.asp
	c := '♬'
	fmt.Printf("The character of c is %c\n", c)
	fmt.Printf("The type of c is %T\n", c)
}
