package sum

import "testing"

func AssertTestPassing(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestProduct(t *testing.T) {

	t.Run("Normal multiplication", func(t *testing.T) {
		got := Product([]int{5, 5, 3})
		want := 75
		AssertTestPassing(got, want, t)
	})

	t.Run("Empty array", func(t *testing.T) {
		got := Product([]int{})
		want := 0
		AssertTestPassing(got, want, t)
	})

	t.Run("One value in array", func(t *testing.T) {
		got := Product([]int{4})
		want := 4
		AssertTestPassing(got, want, t)
	})

	t.Run("Negative numbers", func(t *testing.T) {
		got := Product([]int{4, -2, 5})
		want := -40
		AssertTestPassing(got, want, t)
	})
}

func TestProductAll(t *testing.T) {
	t.Run("Normal multiplication", func(t *testing.T) {
		got := ProductAll([]int{3, 4}, []int{1, 6})
		want := 72
		AssertTestPassing(got, want, t)
	})

	t.Run("Empty array", func(t *testing.T) {
		got := ProductAll([]int{})
		want := 0
		AssertTestPassing(got, want, t)
	})

	t.Run("Multiple empty arrays", func(t *testing.T) {
		got := ProductAll([]int{}, []int{})
		want := 0
		AssertTestPassing(got, want, t)
	})

	t.Run("One empty and one non-empty", func(t *testing.T) {
		got := ProductAll([]int{3, 4}, []int{})
		want := 12
		AssertTestPassing(got, want, t)
	})

	t.Run("Empty slice first", func(t *testing.T) {
		got := ProductAll([]int{}, []int{3, 4})
		want := 12
		AssertTestPassing(got, want, t)
	})
}
