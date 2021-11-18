package main

import "fmt"

func main() {
	var i int

	for { // infinite loop
		if i == 5 { // with this condition => it's not an infinite loop anymore
			break
		}
		fmt.Println(i)
		i++
	}

}
