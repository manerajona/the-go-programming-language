package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10) //1024
	mb                    //= 1024 * kb
	gb                    //= 1024 * mb
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
