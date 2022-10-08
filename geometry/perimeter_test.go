package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{20.0, 5.0}
		checkPerimeter(t, rectangle, 50.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{34.03}
		checkPerimeter(t, circle, 213.81679600332131)
	})
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	/* Table testing is a common approach in
	Golang, we could also iterate over a struct
	and use the keys as the test names: https://github.com/golang/go/wiki/TableDrivenTests */
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 3.0}, want: 30.0},
		{name: "Circle", shape: Circle{Radius: 3.0}, want: 28.274333882308138},
		{
			name:  "Triangle",
			shape: Triangle{Sides: 1, Base: 6},
			want:  36.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.want
			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})
	}
}
