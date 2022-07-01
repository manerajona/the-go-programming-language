package main

import "fmt"

/* OPERATORS:
--------------------
&&	AND
||	OR
!	NEGATE
==	EQUALS
!=	NOT EQUALS
>	GREATER THAN
>=	GREATER OR EQUALS THAN
<	LESS THAN
<=	LESS OR EQUALS THAN
--------------------
*/

var flag bool

func main() {

	fmt.Println("Default value of bool is", flag)

	flag = true
	fmt.Println("New value of bool is", flag)

	// Logical operators
	boo := false
	fmt.Println("AND")
	fmt.Println("Operation flag && boo\t=", flag && boo)
	fmt.Println("Operation flag && true\t=", flag && true)
	fmt.Println("Operation boo && false\t=", boo && false)
	fmt.Println("OR")
	fmt.Println("Operation flag || boo\t=", flag || boo)
	fmt.Println("Operation flag || true\t=", flag || true)
	fmt.Println("Operation boo || false\t=", boo || false)
	fmt.Println("Negate")
	fmt.Println("Operation !flag\t=", !flag)
	fmt.Println("Operation !boo\t=", !boo)

	// Relational operators
	fmt.Println("Logic Op.")
	fmt.Println("Operation 1 == 2\t=", 1 == 2)
	fmt.Println("Operation 1 == 1\t=", 1 == 1)
	fmt.Println("Operation 1 != 2\t=", 1 != 2)
	fmt.Println("Operation 1 != 1\t=", 1 != 1)
	fmt.Println("Operation 1 > 2 \t=", 1 > 2)
	fmt.Println("Operation 1 >= 2\t=", 1 >= 2)
	fmt.Println("Operation 1 < 2 \t=", 1 < 2)
	fmt.Println("Operation 1 <= 2\t=", 1 <= 2)
}
