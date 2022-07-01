package main

import "fmt"

func main() {
	r := sum(2, 3)
	printInt(r)

	x, y := g(r)
	printInt(x)
	printInt(y)

}

// func identifier(params) (returns) { }
func sum(a int, b int) int {
	return a + b
}

// result is pass by VALUE
func printInt(result int) {
	fmt.Printf("result=%d\n", result)
}

// Multiple return
func g(v int) (x, y int) {
	x = v * v
	y = v * v * v
	return // notice that we are not returning any value
}
