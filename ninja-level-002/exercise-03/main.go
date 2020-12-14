package main

import "fmt"

const (
	a        = 3
	b string = "vololo"
)

func main() {
	fmt.Printf("%v\t%T\n%v\t%T", a, a, b, b)
}
