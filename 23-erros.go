package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("arquivo.invalido")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
}
