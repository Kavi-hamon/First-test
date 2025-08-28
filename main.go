package main

import "fmt"

type test interface {
	area() int
}

type circle struct {
	radius int
}

func (c circle) area() int {
	return (c.radius * c.radius * 3)
}

type square struct {
	side int
}

func (s square) area() int {
	return (s.side * s.side)
}
func main() {
	var call test

	a := circle{radius: 10}

	call = a
	fmt.Println(call.area())

	b:= square{side: 10}

	call = b
	fmt.Println(call.area())
}
