package main

// go get code.google.com/p/go-uuid/uuid
import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"log"
)

func main() {
	log.Println("Iniciando o programa")
	fmt.Println("Hello World")
	fmt.Println(uuid.NewRandom())
}
