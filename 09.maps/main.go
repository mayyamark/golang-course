package main

import "fmt"

func main() {
	m := map[string]int{"foo": 42}
	fmt.Println(m) // map[foo:42]

	m["foo"] = 27
	fmt.Println(m) // mutated => map[foo:27]

	delete(m, "foo")
	fmt.Println(m)      // empty => map[]
	fmt.Println(len(m)) // 0

	m["bar"] = 11
	fmt.Println(m)      // map[bar:11]
	fmt.Println(len(m)) // 1
}
