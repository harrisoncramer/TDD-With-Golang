package main

import "fmt"

func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return fmt.Sprintf("Hello, %s", s)
}
