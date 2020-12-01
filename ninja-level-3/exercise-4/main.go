package main

import "fmt"

func main() {
	year := 1950
	for {
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}
