package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Fullname: %v %v\tage: %v\n", p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Pirx",
		last:  "the Pilot",
		age:   33,
	}

	p.speak()
}
