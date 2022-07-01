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

func (e employee) work() {
	fmt.Println(e.firstName, e.lastName, "is working...")
}

func (e *employee) salaryRaise(increment float64) {
	e.salary *= 1. + increment
}

func main() {

	emp := employee{
		person: person{
			firstName: "Steven",
			lastName:  "King",
			age:       50,
		},
		position: "CEO",
		salary:   1_000_000.,
	}
	emp.work()

	fmt.Printf("%s %s salary is %.2f\n", emp.firstName, emp.lastName, emp.salary)
	emp.salaryRaise(.3)
	fmt.Printf("%s %s new salary is %.2f\n", emp.firstName, emp.lastName, emp.salary)

}
