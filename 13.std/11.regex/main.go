package main

import (
	"fmt"
	"regexp"
)

func main() {
	sample := "This is a sample string, theres many strings like this! string, string, string!"

	r, _ := regexp.Compile(`s([a-z]+)g`)
	fmt.Println(r.MatchString(sample))
	fmt.Println(r.FindAllString(sample, -1))
	fmt.Println(r.ReplaceAllString(sample, "text"))
}
