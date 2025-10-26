package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Year", now.Year())
	fmt.Println("Month", now.Month())
	fmt.Println("Day", now.Day())

	// Use layout based on reference time "Mon Jan 2 15:04:05 -0700 MST 2006"
	layout := "02/01/2006"
	fmt.Println("Today is", now.Format(layout))

	future := now.AddDate(10, 0, 0)
	fmt.Println("In 10 years will be", future.Format(layout))

	startDate := time.Date(1989, 4, 4, 3, 45, 0, 0, time.UTC)
	elapsed := now.Sub(startDate)
	fmt.Printf("Hours: %0.f, Minutes: %0.f\n", elapsed.Hours(), elapsed.Minutes())

	elapsed = time.Since(now)
	fmt.Printf("This program took %vms.\n", elapsed.Microseconds())
}
