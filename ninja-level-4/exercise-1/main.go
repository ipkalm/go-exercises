package main

import "fmt"

func main() {
	x := [5]int{1, 3, 5, 7, 9}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T", x)
}
