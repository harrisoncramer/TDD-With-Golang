package geometry

import "testing"

func AssertTestPassing(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}

func AssertTestPassingPrecise(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 5.0}
	got := Perimeter(rectangle)
	want := 50.0
	AssertTestPassing(got, want, t)
}

func TestArea(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 3.0}
		got := rectangle.Area()
		want := 30.0
		AssertTestPassing(got, want, t)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{
			Radius: 3.0,
		}
		got := circle.Area()
		want := 28.274333882308138
		AssertTestPassingPrecise(got, want, t)
	})
}
