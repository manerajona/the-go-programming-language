package main

import (
	"fmt"
	"io"
	"os"
)

type WriteFileError struct {
	Op  string
	err error
}

func (w WriteFileError) Error() string {
	return w.err.Error()
}

func (w WriteFileError) Unwrap() error {
	return w.err
}

type WriteFile struct {
	f   *os.File
	err error
}

func NewWriteFile(filename string) *WriteFile {
	f, err := os.Create(filename)
	if err != nil {
		return &WriteFile{
			f: nil,
			err: WriteFileError{
				Op:  "NewWriteFile",
				err: fmt.Errorf("error while creating a file: %w", err),
			},
		}
	}
	return &WriteFile{
		f:   f,
		err: nil,
	}
}

func (w *WriteFile) WriteString(text string) {
	if w.err != nil {
		return
	}

	if _, err := io.WriteString(w.f, text); err != nil {
		w.err = WriteFileError{
			Op:  "WriteString",
			err: fmt.Errorf("failed while writing a string: %w", err),
		}
	}
}

func (w *WriteFile) Close() {
	if w.f == nil {
		return
	}

	if err := w.f.Close(); err != nil {
		w.err = WriteFileError{
			Op:  "Close",
			err: fmt.Errorf("failed while closing file: %w", err),
		}
	}
}

// Err returns all errors returning from *WriteFile
func (w *WriteFile) Err() error {
	return w.err
}

func main() {
	f := NewWriteFile("file.txt")
	f.WriteString("Hello World")
	f.WriteString("More Text!")
	f.Close()

	// Checking errors
	if err := f.Err(); err != nil {
		panic(err)
	}
	println("success!")
}
