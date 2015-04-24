package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Allisson"] = 32
	m["Chuck Norris"] = 75

	for k, v := range m {
		fmt.Println("Chave:", k, "Valor:", v)
	}

}
