package main

import "fmt"

var x int

func main() {

	// switch case
	x = 999
	fmt.Println("first switch:")
	switch {
	case x > 1000:
		fmt.Println("x>1000")
	case x > 100:
		fmt.Println("x>100")
	case x > 10:
		fmt.Println("x>10")
	}

	// switch value case
	x = 9999
	fmt.Println("second switch:")
	switch x {
	case 9:
		println("nine")
	case 99:
		println("ninety-nine")
	case 999:
		println("nine hundred and ninety-nine")
	case 9999:
		println("nine thousand nine hundred and ninety-nine")
	}

	// fallthrough (opposite to break)
	x = 9999
	fmt.Println("third switch:")
	switch {
	case x > 1000:
		fmt.Println("x>1000")
		fallthrough
	case x > 100:
		fmt.Println("x>100")
		fallthrough
	case x > 10:
		fmt.Println("x>10")
	}

	// default case
	x = 9
	fmt.Println("fourth switch:")
	switch {
	case x > 1000:
		fmt.Println("x>1000")
		fallthrough
	case x > 100:
		fmt.Println("x>100")
		fallthrough
	case x > 10:
		fmt.Println("x>10")
	default:
		fmt.Println("x<10")
	}

	// switch on type
	switch interface{}(x).(type) {
	case int:
		fmt.Println("number")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	default:
		fmt.Println("unknown")
	}

}
