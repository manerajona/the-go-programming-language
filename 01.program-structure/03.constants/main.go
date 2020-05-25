package main

import "fmt"

const pi = 3.14159

const (
	// A represents Van der Pauw constant
	A = 4.53236
	// B represents Kneser-Mahler polynomial constant
	B = 1.38135
)

func main() {

	// Strings
	const (
		s1 = "hello"
		s2
		s3 = "world!"
		s4
	)

	// Booleans
	const TRUE, FALSE = true, false

	// Iota (constant generator)
	const (
		USD int = iota
		EUR
		ARS
		BIT
	)

	fmt.Println("my constants:", pi, A, B, s1, s3, s2, s4, TRUE, FALSE, USD, EUR, ARS, BIT)
}
