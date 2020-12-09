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

	fmt.Printf("%v %v\n", ps1.firstname, ps1.lastname)
	for _, v := range ps1.favIceCreamFlav {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Printf("%v %v\n", ps2.firstname, ps2.lastname)
	for _, v := range ps2.favIceCreamFlav {
		fmt.Printf("\t%v\n", v)
	}
}
