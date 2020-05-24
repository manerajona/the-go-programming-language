package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
	profession          string
}

type celebrity struct {
	person   // unqualified
	knownFor []string
}

func main() {
	var p person
	fmt.Println("The default value for person is", p)

	// Create new person
	p = person{
		firstName:  "Ken",
		lastName:   "Thompson",
		age:        77,
		profession: "Computer scientist",
	}
	fmt.Println("Full name:", p.firstName, p.lastName)
	fmt.Println("Age:", p.age)
	fmt.Println("Job:", p.profession)

	// Composition with structs
	c := celebrity{
		person:   p,
		knownFor: []string{"C", "Unix", "UTF-8", "Go"},
	}
	fmt.Println(c.firstName, c.lastName, "is known for", c.knownFor)

	// Infer fields
	foo := person{
		"John",
		"Smith",
		30,
		"CEO",
	}
	fmt.Println(foo.firstName, foo.lastName, ",", foo.age, ",", foo.profession)

	// Anonymous structs (for use just once)
	message := struct {
		txt  string
		from person
	}{
		"Destroy before reading",
		foo,
	}
	fmt.Println("New message from", message.from.firstName, message.from.lastName, ":", message.txt)

}
