package main

import (
	"fmt"
)

// interface is a collection of method signatures. A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.
// If a type has methods with the same names and signatures as the interface, then it implements that interface.
// write interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// write struct for rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// write method for rectangle to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// write method for rectangle to calculate perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// write struct for circle
type Circle struct {
	Radius float64
}

// write method for circle to calculate area
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// write method for circle to calculate perimeter
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// multiple interface
func (e email) cost() int {
	if !e.isSubscribed {
		return 0
	}
	return 10

}

func (e email) format() string {
	if !e.isSubscribed {
		return "You are not subscribed to our newsletter."
	}
	return e.body

}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

//Type Assertions in Go
//When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value.
// You can cast an interface to its underlying type using a type assertion.

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {

	return 3.14 * c.radius * c.radius
}

func printShapeInfo(s shape) {
	c, ok := s.(circle)
	if ok {
		radius := c.radius
		fmt.Printf("s is a circle, radius: %v\n", radius)
		return
	}
}

func main() {
	var s Shape

	s = Rectangle{Width: 5, Height: 10}
	fmt.Printf("Rectangle Area: %.2f\n", s.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", s.Perimeter())

	s = Circle{Radius: 7}
	fmt.Printf("Circle Area: %.2f\n", s.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", s.Perimeter())
}
