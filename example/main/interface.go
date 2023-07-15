package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	edge float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.edge * r.edge
}
// apparent implementation, no need to declare

func printArea(s Shape) {
	fmt.Println("s.Area() =", s.Area())
}

func main() {
	c := Circle{radius: 1}
	printArea(c)

	r := Rectangle{edge: 1}
	
	var s Shape = r

	str := "Hello World"
	describe(str)			// empty interface, equals to void*

	fmt.Println(s.(Circle)) // casting error

	if circle, ok := s.(Circle); ok { // try type casting
		fmt.Println("s is a circle, r = ", circle.radius)
	}
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}