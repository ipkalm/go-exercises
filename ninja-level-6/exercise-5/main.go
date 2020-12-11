package main

import (
	"fmt"
	"math"
)

type square struct {
	w float64
	l float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.w * s.l
}

func (c circle) area() float64 {
	return math.Pi * c.r * 2
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	var answer float64
	switch s.(type) {
	case square:
		answer = s.(square).area()
	case circle:
		answer = s.(circle).area()
	}
	return answer
}

func main() {
	c1 := circle{
		r: 200.0,
	}

	s1 := square{
		w: 10.0,
		l: 20.0,
	}

	fmt.Println(info(c1))
	fmt.Println(info(s1))
}
