package main

import "fmt"

var (
	nome, sobrenome = "Allisson", "Azevedo"
)

func main() {
	fmt.Println("Olá " + nome + " " + sobrenome)
	// retorna Olá Allisson Azevedo
}
