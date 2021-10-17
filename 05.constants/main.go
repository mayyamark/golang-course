package main

import "fmt"

func main() {
	const pi = 3.14 // the value should be available on COMPILE time (not on run time) => can't be the result of a function
	fmt.Println(pi) // 3.14
	// pi = 3.142 => error

	const a = 3          // implicit type constant
	fmt.Println(a + 3)   // 6
	fmt.Println(a + 1.2) // 4.2 => no error

	const b int = 1
	fmt.Println(b + 3) // 4
	// fmt.Println(b + 1.2) //  error
}
