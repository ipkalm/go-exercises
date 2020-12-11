package main

import "fmt"

func main() {
	f := func(x string) {
		fmt.Println(x)
	}
	f("hello, variable func")
}
