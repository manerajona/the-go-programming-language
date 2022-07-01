package main

import (
	"errors"
	"fmt"
)

var ErrNotExist = fmt.Errorf("the file does not exist")

func openFile(filename string) (string, error) {
	return "", ErrNotExist
}

func main() {
	_, err := openFile("file.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("was unable to open file due %w", err)
		if errors.Is(wrappedErr, ErrNotExist) {
			fmt.Printf("ErrNotExist: %s\n", wrappedErr)
		} else {
			fmt.Printf("Unexpected error: %s\n", wrappedErr)
		}
	}
}
