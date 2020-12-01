package main

import "fmt"

func main() {
	k := 1
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", k)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		k++
	}
}
