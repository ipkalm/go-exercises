package main

import (
	"fmt"

	"github.com/ipkalm/go-exercises/ninja-level-012/exercise-01/dog"
)

func main() {
	hy := 40
	dy := dog.Years(hy)
	fmt.Printf("human years: %v\ndog years: %v\n", hy, dy)
}
