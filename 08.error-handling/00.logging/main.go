package main

import "log"

func main() {
	log.Println("Some util info")
	log.Panicln("Panic!!")         // You can recover of panic
	log.Fatal("Something happend") // You can't recover of fatal
	log.Println("Never gets here")
}
