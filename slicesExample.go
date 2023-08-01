package main

import "fmt"

func main() {

	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x = append(x, 11, 12, 13, 14)
	y := x[2:5]
	fmt.Println(x)
	fmt.Println(y)

}
