package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	t := triangle{height: 5, base: 2}
	s := square{sideLength: 3}
	printArea(t)
	printArea(s)
}

func printArea(a shape) {
	fmt.Println(a.getArea())
}

func (t triangle) getArea() float64 {
	return t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
