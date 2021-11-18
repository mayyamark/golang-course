package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println("--------------------")

	for index, sliceValue := range slice {
		fmt.Println(index, sliceValue)
	}
	fmt.Println("--------------------")

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for key, mapValue := range wellKnownPorts {
		fmt.Println(key, mapValue)

	}
}
