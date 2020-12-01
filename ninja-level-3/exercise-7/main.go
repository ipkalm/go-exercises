package main

import "fmt"

func main() {
	x := "time machine"
	if x == "time machine" {
		fmt.Println(x)
	} else if x == "machine time" {
		fmt.Println(x)
	} else {
		fmt.Println("time error")
	}
}
