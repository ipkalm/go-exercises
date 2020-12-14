package main

import "fmt"

type myOwnType int

var x myOwnType
var y int

func main() {
	fmt.Printf("value: %v\ttype: %T\n", x, x)
	x = 42
	fmt.Println("x = ", x)
	y = int(x)
	fmt.Printf("value: %v\ttype: %T\n", y, y)
}
