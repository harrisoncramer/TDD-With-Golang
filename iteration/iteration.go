package iteration

import (
	"regexp"
	"strings"
)

func Iterate(s string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}

func Uppercase(s string) string {
	firstChar := s[0:1]
	return strings.ToUpper(firstChar) + s[1:]
}

func Words(s string) []string {
	hasSpaces := strings.Contains(s, " ")
	if !hasSpaces {
		if s == "" {
			return []string{}
		} else {
			return []string{s}
		}
	}

	match, err := regexp.MatchString("\\w+", s)
	if err != nil {
		panic("Regular expression could not compile")
	}

	if !match {
		return []string{}
	}

	return strings.Split(s, " ")
}
