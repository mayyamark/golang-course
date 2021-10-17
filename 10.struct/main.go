package main

import "fmt"

func main() {
	type user struct {
		ID   int
		Name string
	}

	var u1 user
	fmt.Println(u1) // empty user struct, with zero-values => { 0 "" }

	u2 := user{ID: 1, Name: "Mayya"}
	fmt.Println(u2) // { 1 "Mayya" }

	u3 := user{
		ID:   2,
		Name: "Mark", // !end with a comma!
	}
	fmt.Println(u3) // { 2 "Mark" }

}
