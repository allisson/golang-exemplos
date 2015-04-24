package main

import "fmt"

func main() {
	alunos := make([]string, 3)
	fmt.Println(alunos)

	alunos[0] = "Fulano"
	alunos[1] = "Beltrano"
	alunos[2] = "Cicrano"
	fmt.Println(alunos)

	alunos = append(alunos, "Zezinho")
	fmt.Println(alunos)

	fmt.Println(alunos[2:4])
}
