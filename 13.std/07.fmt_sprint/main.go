package main

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

type LogLevel int

// color returns the ANSI color code for the level.
func (l LogLevel) color() string {
	switch l {
	case INFO:
		return colorGreen
	case WARN:
		return colorYellow
	case ERROR:
		return colorRed
	default:
		return colorReset
	}
}

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

func main() {
	log(INFO, "Application healthy")
	log(WARN, "Cache miss for key 'user:1234'")
	log(ERROR, "Failed to connect to database")
}

// log builds the colored message using fmt.Sprintf and prints it.
// Accepts only LogLevel, enforcing allowed values at the callsite.
func log(level LogLevel, msg string) {
	out := fmt.Sprintf("%s%s %s", level.color(), msg, colorReset)
	fmt.Println(out)
}
