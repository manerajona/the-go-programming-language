package main

import (
	"errors"
	"log"
)

func main() {
	// Error ignored
	result, _ := sqrt(5)
	log.Println("result of sqrt:", result)

	// Error handled
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err) // program exits here
	}

}

func sqrt(f float64) (result float64, err error) {
	if f < 0 {
		result = 0
		err = errors.New("norgate math: square root of negative number")
	} else {
		result = f * f
		err = nil
	}
	return
}
