package main

import (
	"errors"
	"fmt"
	"log"
)

func Dividir(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Proibida divisão por zero.")
	}
	return x / y, nil
}

func main() {
	valor, err := Dividir(10, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(valor)
}
