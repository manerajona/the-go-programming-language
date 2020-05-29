package main

import (
	"log"
	"os"
)

func main() {
	// Open or create a file 'console.log' where text is appended
	f, _ := os.OpenFile("console.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// NOTE: ignoring errors
	defer f.Close()
	log.SetOutput(f)

	log.Println("Some util info")
	log.Panicln("Panic!!")         // You can recover of panic
	log.Fatal("Something happend") // You can't recover of fatal
	log.Println("Never gets here")
}
