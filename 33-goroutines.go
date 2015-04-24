package main

import (
	"fmt"
	"time"
)

func digaoi(s string) {
	fmt.Println("Dizendo oi de:", s)
}

func main() {
	go digaoi("Goroutine")
	digaoi("Normal")
	time.Sleep(100 * time.Millisecond)
}
