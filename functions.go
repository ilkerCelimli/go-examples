package main

import "fmt"

func main() {
	x := "Resting"
	y := "working"

	if x == "Resting" {
		rest()
	} else if y == "working" {
		work()
	}
}

func rest() {
	fmt.Printf("Resting")
}

func work() {
	fmt.Printf("Work")
}
