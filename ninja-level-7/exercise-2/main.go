package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Pirx",
		last:  "the Pilot",
		age:   33,
	}

	fmt.Println(p.first, p.last, p.age)
	changeMe(&p)
	fmt.Println(p.first, p.last, p.age)
}

func changeMe(p *person) {
	p.first = "Johny"
	p.last = "Mnemonic"
	p.age = 65536
}
