package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * (c.Radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func printShape(s interface{}) {
	switch shape := s.(type) {
	case Rectangle:
		fmt.Println("Length : ", shape.Length, " Width : ", shape.Width)
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())

	case Circle:
		fmt.Println("Radius : ", shape.Radius)
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())

	case Triangle:
		fmt.Println("First side : ", shape.SideA, " Second Side : ", shape.SideB, " Third second : ", shape.SideC)
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())
	default:
		fmt.Println("Provided value does not implement the Shape interface.")
	}
}

func main() {
	rect := Rectangle{Width: 10, Length: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{SideA: 3, SideB: 4, SideC: 5}

	printShape(rect)
	printShape(circle)
	printShape(triangle)

}
