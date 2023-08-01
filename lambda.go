package main

import "fmt"

func main() {
	example := func(x int) bool {
		return x > 0
	}
	y := example(5)
	fmt.Println(y)
}
