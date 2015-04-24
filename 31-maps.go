package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Teste"] = 30
	fmt.Println("O valor é:", m["Teste"])

	m["Teste"] = 31
	fmt.Println("O valor é:", m["Teste"])

	delete(m, "Teste")
	fmt.Println("O valor é:", m["Teste"])

	_, ok := m["Teste"]
	if ok == false {
		fmt.Println("O valor Teste não existe")
	}

}
