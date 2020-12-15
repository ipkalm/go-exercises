package main

import "fmt"

func main() {
	c := make(chan int)
	putNumbers(c)
	pullNumbers(c)
}

func putNumbers(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- (i * 2)
		}
		close(c)
	}()
}

func pullNumbers(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
