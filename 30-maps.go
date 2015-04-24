package main

import "fmt"

func main() {
	m := map[string]int{
		"Allisson":     32,
		"Chuck Norris": 75,
	}

	fmt.Println(m["Allisson"])
	fmt.Println(m["Chuck Norris"])
}
