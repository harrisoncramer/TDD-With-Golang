package dictionary

import (
	"errors"
	"fmt"
)

var (
	ErrWordNotFound = errors.New("Word not found")
	ErrWordExists   = func(word string, def string) error {
		return errors.New(fmt.Sprintf("%q is already defined as %q", word, def))
	}
	ErrDictUndefined = errors.New("Dictionary is not defined")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (definition string, err error) {
	result := d[word]
	if result == "" {
		return "", ErrWordNotFound
	}

	return result, nil
}

/*
We don't have to pass a pointer to modify
the underlying data structure with maps.
This is because, unlike with structs,
maps are basically pointers.
*/
func (dict Dictionary) Define(word string, definition string) (err error) {
	if dict == nil {
		return ErrDictUndefined
	}

	oldDefinition, err := dict.Search(word)
	if err == nil {
		return ErrWordExists(word, oldDefinition)
	}

	dict[word] = definition
	return nil
}

func (dict Dictionary) Update(word string, definition string) (newDefinition string, err error) {
	_, searchError := dict.Search(word)
	if searchError != nil {
		return "", searchError
	}

	dict[word] = definition
	return definition, nil
}
