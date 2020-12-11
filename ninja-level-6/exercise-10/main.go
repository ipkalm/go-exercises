package main

import "fmt"

func main() {
	y := 3
	f := double()

	fmt.Println(f(y))
	fmt.Println(f(y))
	fmt.Println(f(y))
	fmt.Println(f(y))
	fmt.Println(f(y))
	fmt.Println(f(y))

	g := double()
	fmt.Println(g(y))
	fmt.Println(g(y))
	fmt.Println(g(y))
	fmt.Println(g(y))
	fmt.Println(g(y))
}

func double() func(x int) int {
	y := 2
	return func(x int) int {
		y++
		return x * y
	}
}
