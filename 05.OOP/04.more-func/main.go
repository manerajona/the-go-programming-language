package main

import (
	"fmt"
	"time"
)

func main() {

	// Defer func
	defer defered()
	f()

	// Variadic func
	sum := variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("result=%d\n", sum)

	concat := variadicConcat(4, 5., "(", "cuarenta", " y ", "cinco", ")", true)
	fmt.Printf("result=%s\n", concat)

	// Anonymous func
	pow2 := func(num int) int {
		return num * num
	}(sum)
	fmt.Println(sum, "to the power 2 is", pow2)

	g := func(num int) int {
		return num * num * num
	}

	pow3 := g(sum)
	fmt.Println(sum, "to the power 3 is", pow3)

	// Callback
	pow4 := callbackPow(sum, g)
	fmt.Println(sum, "to the power 4 is", pow4)

	// Recursion
	pow5 := recursivePow(sum, 5)
	fmt.Println(sum, "to the power 5 is", pow5)
}

func defered() {
	fmt.Println("Main finished!")
}

func f() {
	for i := 0; i < 100; i++ {
		fmt.Print(".")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Print("\n")
}

func variadicSum(x ...int) (result int) {
	for _, v := range x {
		result += v
	}
	return
}

func variadicConcat(x ...interface{}) (result string) {
	for _, v := range x {
		result += fmt.Sprintf("%v", v)
	}
	return
}

func callbackPow(num int, f func(num int) int) int {
	return num * f(num)
}

func recursivePow(num int, exp int) int {
	if exp == 0 {
		return 1
	}
	exp--
	return num * recursivePow(num, exp)
}
