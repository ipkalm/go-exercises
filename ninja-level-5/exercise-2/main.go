package main

import (
	"fmt"
)

type person struct {
	firstname       string
	lastname        string
	favIceCreamFlav []string
}

func main() {
	ps1 := person{
		firstname: "billy",
		lastname:  "willy",
		favIceCreamFlav: []string{
			"choco",
			"milk",
		},
	}

	ps2 := person{
		firstname: "lerroy",
		lastname:  "jenkins",
		favIceCreamFlav: []string{
			"heroic",
			"courage",
		},
	}

	persons := map[string]person{
		ps1.lastname: ps1,
		ps2.lastname: ps2,
	}

	for k, v := range persons {
		fmt.Printf("\n%v ::\n\t%v likes =>\t", k, v.firstname)
		for _, val := range v.favIceCreamFlav {
			fmt.Printf("%v ", val)
		}
	}
}
