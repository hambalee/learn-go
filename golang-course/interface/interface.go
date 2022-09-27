package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area =", g.area())
	fmt.Println("Perimeter =", g.perimeter())
}

func main() {
	r := rectangle{width: 5, height: 10}
	c := circle{radius: 15}

	measure(r)
	// 	{5 10}
	// Area = 50
	// Perimeter = 30
	measure(c)
	// 	{15}
	// Area = 706.8583470577034
	// Perimeter = 94.24777960769379
}
