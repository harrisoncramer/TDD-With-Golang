package mocking

import (
	"fmt"
	"io"
)

/*
We define an interface which defines the sleep function. This can be used
as our dependency (see the dependency-injection code for a simpler example).
Next, we define a Spy struct that has a Sleep method, which tracks how many
times it's been called.

When this function is actually used, we could pass the time package, which
has a Sleep method. But when testing, we pass in a mock that spys on this
Sleep method.
*/
type Sleeper interface {
	Sleep(time int)
}

func Countdown(b io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(b, fmt.Sprint(i))
		sleeper.Sleep(1)
	}
	fmt.Fprint(b, "go!")
}
