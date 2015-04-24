package main

import "fmt"

type Numero struct {
	X int
}

func (n Numero) Dobro() int {
	n.X = n.X * 2
	return n.X
}

func (n *Numero) Dobro2() int {
	n.X = n.X * 2
	return n.X
}

func main() {
	numero := Numero{X: 10}
	fmt.Println(numero.Dobro())
	fmt.Println(numero.X)

	fmt.Println(numero.Dobro2())
	fmt.Println(numero.X)
}
