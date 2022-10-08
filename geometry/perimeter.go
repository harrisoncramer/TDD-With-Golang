package geometry

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Sides float64
	Base  float64
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

func (t Triangle) Area() float64 {
	return (t.Base * t.Sides) / 2
}

func (c Circle) Perimeter() float64 {
	return (2.0 * c.Radius) * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() float64 {
	return t.Base + (t.Sides * 2)
}
