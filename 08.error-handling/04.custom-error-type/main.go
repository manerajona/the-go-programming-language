package main

import (
	"fmt"
	"log"
)

// MathError refers to square root of negative number error
type MathError struct {
	code string
	info error // implicitly implements error interface
}

func (n *MathError) Error() string {
	return fmt.Sprintf("Math error %v occured: %q", n.code, n.info)
}

func main() {
	_, err := sqrt(-10.01)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (result float64, err error) {
	if f < 0 {
		info := fmt.Errorf("Norgate math redux: square root of negative number: %v", f)
		result, err = 0, &MathError{"001", info}
		return
	}
	result, err = (f * f), nil
	return
}
