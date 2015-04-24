package main

import "fmt"

func main() {
	alunos := make([]string, 3)
	alunos[0] = "Fulano"
	alunos[1] = "Beltrano"
	alunos[2] = "Cicrano"

	for i, value := range alunos {
		fmt.Println("Posição:", i, "Valor:", value)
	}
}
