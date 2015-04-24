package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(s string, channel chan string) {
	res, err := http.Get(s)

	if err != nil {
		log.Println("Não foi possível conectar com a url", s)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	result := fmt.Sprintf("A página: %s, Tem o tamanho de %d", s, len(string(body)))
	channel <- result
}

func main() {
	channel := make(chan string)

	go fetch("http://linuxmint.com/", channel)
	go fetch("http://www.ubuntu.com/", channel)
	go fetch("http://www.debian.org/", channel)
	go fetch("http://www.opensuse.org/en/", channel)
	go fetch("http://getfedora.org/", channel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-channel)
	}
}
