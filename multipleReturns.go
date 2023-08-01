package main

import "fmt"

func main() {
	_, b := multiReturn() // _ dikkate alınmayacak değer anlamına gelir.
	fmt.Println(b)
	ex, y := namedReturn(55)
	fmt.Println(ex, y)
}

func multiReturn() (a string, b string) {
	return "hello", "world"
}

func namedReturn(a int) (x, y int) {
	x = a + 1
	y = 0
	return
}
