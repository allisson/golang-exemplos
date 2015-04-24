package main

import "fmt"

func zera1(valor int) {
	valor = 0
}

func zera2(valor *int) {
	*valor = 0
}

func main() {
	valor := 10

	zera1(valor)
	fmt.Println(valor)
	// retorna 10

	zera2(&valor)
	fmt.Println(valor)
	// retorna 0
}
