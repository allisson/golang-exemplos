package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("teste.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("Teste do defer")
}
