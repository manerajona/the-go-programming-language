package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Deferred RECOVER
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	fmt.Println(readFileContent("file.txt"))
	fmt.Println(readFileContent("no-file.txt"))
}

func readFileContent(path string) string {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		// PANIC!!
		panic(err)
	}
	return string(content)
}
