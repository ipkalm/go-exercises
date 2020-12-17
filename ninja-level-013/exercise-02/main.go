package main

import (
	"fmt"

	"github.com/ipkalm/go-exercises/ninja-level-013/exercise-02/quote"
	"github.com/ipkalm/go-exercises/ninja-level-013/exercise-02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
