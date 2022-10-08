package geometry

import "math"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
