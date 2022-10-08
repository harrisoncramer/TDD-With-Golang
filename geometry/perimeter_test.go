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
		shape Shape
		want  float64
	}{
		{shape: Rectangle{10.0, 3.0}, want: 30.0},
		{shape: Circle{3.0}, want: 28.274333882308138},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		want := test.want
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
}
