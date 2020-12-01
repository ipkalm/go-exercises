package main

import (
	"fmt"
)

func main() {
	switch {
	case (1 == 0):
		fmt.Println("hello false")
	case (1 == 1):
		fmt.Println("hello true")
	}
}
