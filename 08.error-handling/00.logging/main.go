package main

import "log"

func main() {
	log.Println("Some util info")
	log.Fatal("Something happend")
	log.Panicln("Panick!") // Panick kills the program
	log.Println("Never gets here")
}
