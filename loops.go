package main

import "fmt"

func main() {
	x := make([]int, 1)

	for count := 0; count < 999; count++ {
		fmt.Printf("hello world")
		x = append(x, count)
	}

	for _, val := range x {
		fmt.Println(val)
	}
	var f = 5
	var y = 44
	for f != y {
		f++
		fmt.Println(f)
	}
}
