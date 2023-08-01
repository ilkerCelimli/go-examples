package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 4564
	var y = &x
	getPointer(y)
	x = 44444
	getPointer(y)
}

func getPointer(pointer *int) {
	fmt.Println(*pointer)

}

type Person struct {
	name    string
	surname string
	birtday time.Time
}
