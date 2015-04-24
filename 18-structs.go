package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	pessoa := Pessoa{Nome: "Fulano", Idade: 1}
	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)

	allisson := &pessoa
	allisson.Nome = "Allisson"
	allisson.Idade = 32

	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)
}
