package main

import "fmt"

var m map[string]string

func main() {

	fmt.Println("Default value of map:", m)

	// map [key]value{}
	m = map[string]string{
		"Go":    "Google",
		"Swift": "Apple",
		"Java":  "Oracle",
		"C#":    "Microsoft",
	}
	fmt.Println(m["Go"])
	fmt.Println(m["Swift"])
	fmt.Println(m["Java"])
	fmt.Println(m["C#"])
	//fmt.Println(m["Javascript"])

	// Evaluate if value is present
	if value, ok := m["Javascript"]; ok {
		println(value) // never get there
	} else {
		println(`For key "Javascript" value doesn't exists`)
	}

	// Add element
	m["Python"] = "Python Software Foundation"

	// for range
	for k, v := range m {
		fmt.Printf("For key \"%s\" value is \"%s\"\n", k, v)
	}

	// delete
	delete(m, "Python")
	fmt.Println(m)

}
