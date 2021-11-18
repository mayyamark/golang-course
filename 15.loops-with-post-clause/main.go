package main

import "fmt"

func main() {
	var i int

	// Loop till a condition clause with a post a clause
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// fmt.Println(j) // undefined
}
