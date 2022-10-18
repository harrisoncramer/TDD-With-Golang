package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	tests := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Harry"},
			[]string{"Harry"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Harry", "New York"},
			[]string{"Harry", "New York"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Harry", 29},
			[]string{"Harry"},
		},
		{
			"Nested struct",
			Person{"Harry", Profile{29, "New York"}},
			[]string{"Harry", "New York"},
		},
		{
			"Pointer to a struct",
			&Person{
				"Harry",
				Profile{
					29,
					"New York",
				},
			},
			[]string{"Harry", "New York"},
		},
		{
			"Slice of structs",
			[]Profile{
				{22, "London"},
				{30, "Amsterdam"},
			},
			[]string{"London", "Amsterdam"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"}},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"foo": "bar",
				"baz": "boz",
			},
			[]string{"bar", "boz"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v wanted %v", got, test.ExpectedCalls)
			}
		})
	}

	/* We cannot guarantee order with maps */
	t.Run("Works with maps", func(t *testing.T) {
		testCase := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		got := []string{}

		walk(testCase, func(s string) {
			got = append(got, s)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, val := range haystack {
		if val == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("Expected %+v to contain %q but didn't", haystack, needle)
	}
}
