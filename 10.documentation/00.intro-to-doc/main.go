// Copyright 2022 Jonathan Manera. All rights reserved.

// Package employee implements methods and attributes of an employee
package employee

import "fmt"

// Person refers to a person with Firstname, Lastname and Age
type Person struct {
	FirstName, LastName string
	Age                 int
}

// Employee refers to a person with Person, Position and Salary
type Employee struct {
	Person
	Position string
	Salary   float64
}

// work is a trivial implementation of Employee
func (e Employee) work() {
	fmt.Println(e.FirstName, e.LastName, "is working...")
}

// salaryRaise accept a decimal and rise the salary of Employee
// Example: an increment of 20% is represented by .2
func (e *Employee) salaryRaise(increment float64) {
	e.Salary *= 1. + increment
}

// run '$ go doc Employee' to see the documentation
