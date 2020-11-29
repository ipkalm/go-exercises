package main

import "fmt"

func main() {
	a := 10 == 20
	b := 42 <= 7
	c := 9 >= 23
	d := 77 != 12
	e := 1 < 1
	f := 100 > 2222

	fmt.Println(a, b, c, d, e, f)
}
