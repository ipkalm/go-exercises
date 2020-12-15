package main

import "fmt"

const counter int = 10

func main() {
	ch := make(chan int)

	for i := 0; i < counter; i++ {
		go func() {
			for j := 0; j < counter; j++ {
				ch <- j
			}
		}()
	}

	for k := 0; k < counter*counter; k++ {
		fmt.Println(<-ch)
	}
}
