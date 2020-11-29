package main

import "fmt"

type myOwnType int

var x myOwnType

func main() {
	fmt.Printf("value: %v\ttype: %T\n", x, x)
	x = 42
	fmt.Println(x)
}
