package main

import "fmt"

const Pi = 3.14
const (
	StatusOK       = 200
	StatusNotFound = 404
)

func main() {
	fmt.Println(Pi)
	// retorna 3.14
	fmt.Println(StatusOK)
	// retorna 200
	fmt.Println(StatusNotFound)
	// retorna 404
}
