package main

import "fmt"

func main() {
	z := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	func(func(x ...int) int, []int) {
		fmt.Println(calc(z...))
	}(calc, z)
}

func calc(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
