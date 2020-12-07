package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}
	fmt.Println(xy)

	for _, v := range xy {
		for _, v1 := range v {
			fmt.Printf("%v ", v1)
		}
		fmt.Printf("\n")
	}
}
