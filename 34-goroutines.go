package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetch(s string) {
	res, err := http.Get(s)

	if err != nil {
		log.Println("Não foi possível conectar com a url", s)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("A página:", s, "Tem o tamanho de:", len(string(body)))
}

func main() {
	go fetch("http://www.google.com.br")
	fetch("http://www.uol.com.br")
	time.Sleep(5000 * time.Millisecond)
}
