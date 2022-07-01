package main

import (
	"fmt"
	"log"
)

// ErrNorgateMath refers to square root of negative number error
var ErrNorgateMath = fmt.Errorf("norgate math: square root of negative number")

func main() {
	// errors.errorString
	fmt.Printf("%T\n", ErrNorgateMath)

	result, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Result of sqrt:", result) // never gets here
}

func sqrt(f float64) (result float64, err error) {
	if f < 0 {
		result = 0
		err = ErrNorgateMath
	} else {
		result = f * f
		err = nil
	}
	return
}
