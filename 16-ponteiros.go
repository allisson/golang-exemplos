package main

import "fmt"

func main() {
	a := 10
	b := a
	c := &a

	b = 20
	*c = 30

	fmt.Println(a)
	// retorna 30
	fmt.Println(b)
	// retorna 20
	fmt.Println(*c)
	// retorn 30
}
