package main

import "fmt"

func main() {

	s := "Every value have a pointer"

	fmt.Printf("s value:\t\t%q\n", s)
	fmt.Printf("s type:\t\t\t%T\n", s)

	fmt.Printf("s value memory address:\t%v\n", &s)
	fmt.Printf("s pointer type:\t\t%T\n", &s)

	// Assign pointer
	var p *string = &s
	fmt.Printf("p value:\t\t%v\n", p)
	fmt.Printf("p reference to value:\t%q\n", *p)

	// p is always pointing to s
	s = "A pointer *string is a type"
	fmt.Printf("s value:\t\t%q\n", s)
	fmt.Printf("p reference to value:\t%q\n", *p)

	// Modify s using p
	*p = "A pointer is just an address in memory"
	fmt.Printf("s value:\t\t%q\n", s)
}
