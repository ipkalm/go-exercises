package main

import (
	"fmt"

	"github.com/ipkalm/go-exercises/ninja-level-013/exercise-01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(700))
}
