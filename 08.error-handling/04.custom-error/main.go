package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// MathError refers to square root of negative number error
type MathError struct {
	code   int
	detail string
	at     time.Time
}

// Implements error interface
func (e MathError) Error() string {
	return fmt.Sprintf("Math error %v occured: %q at %v", e.code, e.detail, e.at)
}

func (e MathError) Is(other error) bool {
	_, ok := other.(MathError)
	return ok
}

func main() {
	_, err := sqrt(-10.01)
	if errors.Is(err, MathError{}) {
		log.Println(err)
	}
}

func sqrt(f float64) (result float64, err error) {
	if f < 0 {
		detail := fmt.Sprintf("norgate math redux: square root of negative number: %v", f)
		result, err = 0, &MathError{400, detail, time.Now()}
	} else {
		result = f * f
	}
	return
}
