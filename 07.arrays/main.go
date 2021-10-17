package main

import "fmt"

func main() {
	var arr [3]int // fixed size, all elements have the same type
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	fmt.Println(arr == arr2) // true

}
