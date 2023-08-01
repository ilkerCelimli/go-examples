package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["hello"] = "world"
	fmt.Println(m)
	m["hello"] = "mars"
	fmt.Println(m)
}
