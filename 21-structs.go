package main

import "fmt"

type Pessoa struct {
	Nome string
}

type Funcionario struct {
	Pessoa
	Mat int
}

func main() {
	funcionario := Funcionario{Pessoa{Nome: "Allisson"}, 12345}
	fmt.Println(funcionario.Nome)
	fmt.Println(funcionario.Mat)
}
