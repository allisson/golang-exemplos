package main

import (
	"fmt"
	"strings"
)

func nomeSobrenome(nome string) (string, string) {
	return strings.Split(nome, " ")[0], strings.Split(nome, " ")[1]
}

func main() {
	nome, sobrenome := nomeSobrenome("Allisson Azevedo")
	fmt.Println("Nome = " + nome)
	fmt.Println("Sobrenome = " + sobrenome)
}
