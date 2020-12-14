package main

import "fmt"

type person struct {
	Name string
}

func (p *person) speak() {
	fmt.Println(p.Name)
}

// func (p person) speak() {
// 	fmt.Println(p.Name)
// }

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{"Charlie"}
	saySomething(&p) //can with `func (p person) speak()` declaration
	// saySomething(p)  //can with `func (p person) speak()` declaration

	saySomething(&p) //can with `func (p *person) speak()` declaration
	// saySomething(p)  //CAN'T with `func (p *person) speak()` declaration
}
