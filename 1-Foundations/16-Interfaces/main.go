package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

func showArea(s shape) {
	fmt.Printf("The object area is %0.2f", s.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}


// Generic type
func generic(interf interface{}) {
    fmt.Println(interf)
}

func main() {

	r := rectangle{10, 15}
	showArea(r)

	c := circle{10}
	showArea(c)

    // Generic type
    generic("string")
    generic(1)
    generic(true)


}
