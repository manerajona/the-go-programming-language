package main

import "fmt"

var x []int

func main() {

	fmt.Println("Default value of a slice", x)

	// composite literal
	x = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	fmt.Println(x)

	// subset of a slice
	y := x[0:5]
	fmt.Println(y)
	y = x[:len(x)/2]
	fmt.Println(y)

	y = x[5:10]
	fmt.Println(y)
	y = x[len(x)/2:]
	fmt.Println(y)

	// for each
	for index, item := range x {
		fmt.Println("for index", index, "the value is", item)
	}

	// append (add new item)
	x = append(x, 55)
	x = append(x, 60)
	x = append(x, 65)
	x = append(x, 70, 75, 80, 85, 90, 95, 100)
	fmt.Println(x)

	y = append(y, x...)
	fmt.Println(y)

	// delete on slices
	y = append(y[:5], y[15:]...)
	fmt.Println(y)

	y = y[4:]
	fmt.Println(y)

	// make
	var z = make([]int, 10, 100)
	fmt.Println("Default value of z is", z)
	// The length is the number of slice elements
	fmt.Println("The length of z is", len(z))
	// Slice's "underlying array" is determined by the capacity
	fmt.Println("The capacity of z is", cap(z))
	// The length can not exceed the capacity

}
