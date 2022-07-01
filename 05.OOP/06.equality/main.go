package main

import "fmt"

type person struct {
	id                  int64
	firstName, lastName string
	age                 int32
}

func (p person) String() string {
	return fmt.Sprintf("person=(firstName=%v,lastName=%v,age=%v)", p.firstName, p.lastName, p.age)
}

func (p person) Equals(p0 person) bool {
	if p.id == p0.id {
		return true
	}
	return false
}

func main() {
	p1 := person{id: 1, firstName: "John", lastName: "Doe", age: 41}
	p2 := person{id: 1, firstName: "John", lastName: "Doe", age: 41}
	p3 := person{id: 2, firstName: "Jane", lastName: "Smith", age: 26}
	p4 := person{id: 2}

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)
	fmt.Println(p3 == p4)
	fmt.Println(p3.Equals(p4))
	fmt.Println(p3.Equals(p1))
}
