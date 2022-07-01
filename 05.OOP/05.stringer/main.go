package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func (p person) String() string {
	return fmt.Sprintf("person=(firstName=%v,lastName=%v,age=%v)", p.firstName, p.lastName, p.age)
}

func main() {
	p := person{firstName: "John", lastName: "Doe", age: 41}
	fmt.Println(p)
}
