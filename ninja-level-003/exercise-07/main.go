package main

import "fmt"

func main() {
	x := "machine time"
	if x == "time machine" {
		fmt.Println(x)
	} else if x == "machine time" {
		fmt.Println(x)
	} else {
		fmt.Println("time error")
	}
}
