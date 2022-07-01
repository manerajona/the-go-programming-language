package main

import "fmt"

func main() {

	// for init; condition; postinit {do something...}
	for i := 0; i < 10; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println("done")

	// for condition {do something...}
	j := 0
	for j != 10 {
		fmt.Print(j, ",")
		j++
	}
	fmt.Println("done")

	// break and continue
	for x := 0; ; x++ {
		if x > 100 {
			break
		}
		if x%5 != 0 {
			continue
		}
		fmt.Print(x, ",")
	}
	fmt.Println("done")

	// infinite loop
	y := 10
	for {
		if y == 0 {
			break
		}
		fmt.Print(y, ",")
		y--
	}
	fmt.Println("done")

	// Range loops
	const nihongo = "日本語"
	for _, value := range nihongo {
		fmt.Printf("%c,", value)
	}
	fmt.Println("done")

}
