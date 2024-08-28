package main

import (
	"fmt"
	"math"
)
// func returnType
type geometry interface {
	area() float64
	perim() float64
}

type reactangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r reactangle) area() float64{
	return r.width * r.height
}

func (r reactangle) perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

// use interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area = ", g.area())
	fmt.Println("perim = ", g.perim())
}

func main() {
	r := reactangle{width: 5, height:7}
	c := circle{radius:3}

	measure(r)
	measure(c)
}