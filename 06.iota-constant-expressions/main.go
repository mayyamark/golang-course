package main

import "fmt"

const (
	a1 = 1
	a2 = "asd"
)

const (
	b1 = iota
	b2 = iota
)

const (
	c1 = iota + 6
	c2 = iota
)

const (
	d1 = iota + 6
	d2
)

const (
	f1 = iota + 6
	f2 = 2 << iota
	f3 = 32 >> iota
)

func main() {
	fmt.Println(a1, a2)     // 1	asd
	fmt.Println(b1, b2)     // 0	1
	fmt.Println(c1, c2)     // 6	1
	fmt.Println(d1, d2)     // 6	6+1 => 7
	fmt.Println(f1, f2, f3) // 6	{n times *2} => 2*1=2 (x2) => 4		{/2 n times}   => 8
}
