package main

import "fmt"

func soma(x, y int) int {
	return x + y
}

func main() {
	fmt.Print("2 + 2 = ")
	fmt.Println(soma(2, 2))
}
