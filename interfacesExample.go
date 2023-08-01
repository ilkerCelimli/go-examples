package main

import "fmt"

func main() {
	domates := DomatesSalcası{}
	var salca Salca
	salca = &domates
	salca.Ye()
	biber := BiberSalcası{}
	salca = &biber
	salca.Ye()
}

type Salca interface {
	Ye()
}

type BiberSalcası struct {
}
type DomatesSalcası struct {
}

func (salca BiberSalcası) Ye() {
	fmt.Println("Biber Salcası yeniliyor")
}
func (salca DomatesSalcası) Ye() {
	fmt.Println("Domates Salcası yeniliyor")
}
