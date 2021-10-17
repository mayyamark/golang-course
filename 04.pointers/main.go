package main

import "fmt"

func main() {
	var a *string // points to the address in the memory
	// a = "aaa" => error
	// *a = "aaa" => error
	fmt.Println(a) // nil, empty pointer

	var b *string = new(string)
	*b = "bbb"
	fmt.Println(b, *b) // address	bbb

	dummy := "dummy"
	pointer := &dummy              // create a pointer and give it the value of dummy
	fmt.Println(pointer, *pointer) // address	dummy
}
