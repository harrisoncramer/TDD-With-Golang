package geometry

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

/*
In Golang, type resolution is implicit,
we don't have to "tell" the compiler that
Circle and Rectangle implement this interface
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Perimeter() float64 {
	return (2.0 * c.Radius) * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
