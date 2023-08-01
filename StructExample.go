package main

import (
	"fmt"
	"time"
)

func main() {
	first := example{4}
	second := &first
	fmt.Println(second.age)
	fmt.Println(first.age)
}

type person struct {
	name    string
	surname string
	birtday time.Time
}

type example struct {
	age int
}
