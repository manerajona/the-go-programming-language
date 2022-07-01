package main

import "fmt"

/* VERBS:
----------------------------------------
%v	the value in a default format
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type
%%	a literal percent sign; consumes no value

%t	boolean

%d	decimal integer
%c	rune

%s	string
%q	quoted string "abc"
----------------------------------------
more verbs on: https://golang.org/pkg/fmt/
*/

/* SCAPES:
--------------------
\n	line break
\b	backspace
\t	horizontal tab
\v	vertical tab
\\	backslash
--------------------
*/

func main() {
	// Number
	i := 42

	fmt.Println("my value i =", i)
	fmt.Printf("my value i = %d\n", i)
	fmt.Printf("my value i = %v\n", i)

	fmt.Printf("my value type i is %T\n", i)

	// Boolean
	b := true

	fmt.Println("my value b =", b)
	fmt.Printf("my value b = %t\n", b)
	fmt.Printf("my value b = %v\n", b)

	fmt.Printf("my value type b is %T\n", b)

	// string
	s := "Hello world!"

	fmt.Println("my value s =", s)
	fmt.Printf("my value s = %s\n", s)
	fmt.Printf("my value s = %v\n", s)
	fmt.Printf("my value s quoted is %q\n", s)

	fmt.Printf("my value type s is %T\n", s)

}
