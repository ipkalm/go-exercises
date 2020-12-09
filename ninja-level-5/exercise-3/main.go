package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	spectra := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	ram := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Printf("spectra:\n\tdoors: %v\n\tcolor: %v\n\tluxury?: %v\n", spectra.doors, spectra.color, spectra.luxury)
	fmt.Printf("ram:\n\tdoors: %v\n\tcolor: %v\n\t4x4?: %v", ram.doors, ram.color, ram.fourWheel)
}
