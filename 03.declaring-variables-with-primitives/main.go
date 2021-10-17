package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i) // 42

	var a float32 = 3.14
	fmt.Println(a) // 3.14

	var b float64 = 12.34
	fmt.Println(b) // 12.34

	firstName := "Mayya"
	fmt.Println(firstName) // Mayya

	c := complex(3, 4)
	fmt.Println(c) // (3+4i)

	r, im := real(c), imag(c)
	fmt.Println(r, im) // 3 4

	d := true
	fmt.Println(d) // true
}
