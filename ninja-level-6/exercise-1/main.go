package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x)
	fmt.Println(y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "string 42"
}
