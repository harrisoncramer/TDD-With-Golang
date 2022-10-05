package main

import "fmt"

const englishGreeting = "Hello"
const spanishGreeting = "Hola"
const frenchGreeting = "Bonjour"

func Hello(s string, l string) string {
	if s == "" {
		s = "world"
	}

	g := englishGreeting

	switch l {
	case "Spanish":
		g = spanishGreeting
	case "French":
		g = frenchGreeting
	}

	if l == "Spanish" {
		g = spanishGreeting
	}

	return fmt.Sprintf("%s, %s", g, s)
}
