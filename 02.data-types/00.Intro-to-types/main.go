package main

import "fmt"

type myType string

var s myType

func main() {

	s = "Go is all about types"

	fmt.Printf("s value is %q\n", s)
	fmt.Printf("s type is %T\n", s)

}
