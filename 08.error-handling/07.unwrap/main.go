package main

import (
	"errors"
	"fmt"
)

type ErrFile struct {
	Filename string
	Wrapped  error
}

func (e ErrFile) Error() string {
	return fmt.Sprintf("File %s: %v", e.Filename, e.Wrapped)
}

func (e ErrFile) Unwrap() error {
	return e.Wrapped
}

var ErrNotExist = fmt.Errorf("the file does not exist")

func openFile(filename string) (string, error) {
	return "", ErrFile{
		Filename: filename,
		Wrapped:  ErrNotExist,
	}
}

func main() {
	_, err := openFile("file.txt")
	if err != nil {
		var fErr ErrFile
		if errors.As(err, &fErr) {
			//fmt.Printf("ErrFile: was unable to open file %s due %s\n", fErr.Filename, fErr.Unwrap().Error())
			fmt.Printf("ErrFile: was unable to open file %s due %s\n", fErr.Filename, errors.Unwrap(fErr))
		} else {
			fmt.Printf("Unexpected error: was unable to open file due %s\n", err.Error())
		}
	}
}
