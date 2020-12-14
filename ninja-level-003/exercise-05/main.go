package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		z := i % 4
		if z != 0 {
			fmt.Println(z)
		}
	}
}
