package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

func Greet(w io.Writer, s string) {
	fmt.Fprintf(w, "Hello, %s", s)
}

/* Untestable wrapper that logs to stdout */
func GreetOut(s string) {
	Greet(os.Stdout, s)
}
