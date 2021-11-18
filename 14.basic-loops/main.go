package main

import "fmt"

func main() {
	var i, j, k int
	// for i < 5 {} // Infinite loop

	// Loop till a condition clause
	for i < 5 {
		fmt.Println("Loop till a condition clause: ", i)
		i++
	}

	fmt.Println("-------------------------------------------")

	// Loop till a condition with BREAK
	for j < 5 {
		fmt.Println("Loop till a condition with BREAK: ", j)
		j++

		if j == 3 {
			break
		}
	}

	fmt.Println("-------------------------------------------")

	// Loop till a condition CONTINUE
	for k < 5 {
		fmt.Println("Loop till a condition with CONTINUE: ", k)
		k++

		if k == 3 {
			continue
		}
		fmt.Println("continuing...")

	}
}
