package main

import "fmt"

func main() {
	f := fReturner()
	f()

	fReturner()()
}

func fReturner() func() {
	return func() {
		fmt.Println("Returned func")
	}
}
