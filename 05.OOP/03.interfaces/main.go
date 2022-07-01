package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

type employee struct {
	person
	position string
	salary   float64
}

/*
 * POLYMORPHISM
 */
func (e employee) work() {
	fmt.Println(e.firstName, e.lastName, "is working...")
}

func (p person) work() {
	fmt.Println("No working,", p.firstName, "resting!")
}

type human interface {
	work()
}

func main() {

	steven := person{
		firstName: "Steven",
		lastName:  "King",
		age:       50,
	}

	emp := employee{
		person:   steven,
		position: "CEO",
		salary:   100_000_000.,
	}

	giveOrder(emp)
	giveOrder(steven)

	who(emp)
	who(steven)
	who(struct{ human }{})

}

func giveOrder(h human) {
	h.work()
}

func who(h human) {
	switch h.(type) {
	case person:
		fmt.Println("This is a person called", h.(person).firstName, ", he's", h.(person).age, "years old.")
	case employee:
		fmt.Println("This is an employee called", h.(employee).firstName, h.(employee).lastName, ", he's a", h.(employee).position, ".")
	default:
		fmt.Println("This is a human being.")
	}
}
