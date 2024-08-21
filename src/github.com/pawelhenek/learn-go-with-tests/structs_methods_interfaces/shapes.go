package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// methods

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
	return triangle.Height * triangle.Base / 2
}

func main() {
	var f1 = Circle{10}
	var f2 = Rectangle{5, 10}
	var f3 = Triangle{7, 6}

	fmt.Println(fmt.Sprintf("Circle Area: %.2f", f1.Area()))
	fmt.Println(fmt.Sprintf("Rectangle Area: %.2f", f2.Area()))
	fmt.Println(fmt.Sprintf("Triangle Area: %.2f", f3.Area()))
}
