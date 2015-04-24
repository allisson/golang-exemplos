package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Allisson"] = 32
	m["Chuck Norris"] = 75

	fmt.Println(m["Allisson"])
	fmt.Println(m["Chuck Norris"])
}
