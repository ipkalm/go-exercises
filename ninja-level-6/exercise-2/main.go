package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(foo(1, 2, 3))
	fmt.Println(bar(x))
}

func foo(x ...int) int {
	var z []int

	for _, v := range x {
		z = append(z, v)
	}

	sum := 0
	for _, v := range z {
		sum += v
	}

	return sum
}

func bar(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
