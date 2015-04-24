package main

import (
	"fmt"
	"strings"
)

func nomeSobrenome(nomeCompleto string) (nome, sobrenome string) {
	nome = strings.Split(nomeCompleto, " ")[0]
	sobrenome = strings.Split(nomeCompleto, " ")[1]
	return
}

func main() {
	nome, sobrenome := nomeSobrenome("Allisson Azevedo")
	fmt.Println("Nome = " + nome)
	fmt.Println("Sobrenome = " + sobrenome)
}
