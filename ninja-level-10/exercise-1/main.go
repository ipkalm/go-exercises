package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	bc := make(chan int, 1)

	go func() { c <- 42 }()
	bc <- 43

	fmt.Println(<-c)
	fmt.Println(<-bc)
}
