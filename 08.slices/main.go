package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}

	slice := arr[:]         // create a slice, that contains the values of the arr array => get slice, but with fixed size again
	fmt.Println(arr, slice) // [1 2 3]	[1 2 3]
	fmt.Println(arr, slice) // [1 2 3]	[1 2 3 3]

	// fmt.Println(arr == slice) // error => mismatched types

	arr[0] = 11
	slice[1] = 22
	fmt.Println(arr, slice) // both are mutated => [11 22 3]	[11 22 3]

	slice2 := []int{1, 2, 3, 4}
	slice2 = append(slice2, 5)
	fmt.Println(slice2) // [1 2 3 4 5]

	slice3 := append(slice2, 6, 7)
	slice3[0] = 0
	fmt.Println(slice2, slice3) // both are mutated => [0 2 3 4 5] [0 2 3 4 5 6 7]

	s1 := slice3[1:]        // start from index 1
	s2 := slice3[:2]        // to index 2 (not included)
	s3 := slice3[1:3]       // from index 1 to index 3
	fmt.Println(s1, s2, s3) // [2 3 4 5 6 7]	[0 2]	[2 3]
}
