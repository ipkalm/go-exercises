package main

import "fmt"

func main() {
	x := struct {
		name     string
		visited  map[string]bool
		wishList []string
	}{
		name: "Val",
		visited: map[string]bool{
			"France":  true,
			"Belgium": true,
			"Poland":  false,
			"Russia":  true,
		},
		wishList: []string{
			"USA",
			"Australia",
			"Thailand",
		},
	}

	fmt.Printf("The name: %v\n\tvisited countries:\n", x.name)
	for k, v := range x.visited {
		if v {
			fmt.Printf("\t\t%v\n", k)
		}
	}
	fmt.Printf("\tcoutries wish to visit:\n")
	for _, v := range x.wishList {
		fmt.Printf("\t\t%v\n", v)
	}
}
