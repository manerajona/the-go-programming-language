package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")

	var pathError *os.PathError

	switch {
	case errors.Is(err, os.ErrPermission) && errors.As(err, &pathError):
		e := fmt.Errorf("you do not have permission to open the file: %w and the path is %s", err, pathError.Path)
		log.Println(e)
	case errors.Is(err, os.ErrNotExist) && errors.As(err, &pathError):
		e := fmt.Errorf("the file does not exist: %w and the path is %s", err, pathError.Path)
		log.Println(e)
	case errors.As(err, &pathError):
		e := fmt.Errorf("here is the orignial error %w and the path %s", err, pathError.Path)
		log.Println(e)
	case err != nil:
		e := fmt.Errorf("file couldn't be opened: %w", err)
		log.Println(e)
	}

	defer f.Close()
}
