package maps

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// ?? map-lookup returns 2 values.
	// ?? the first one is the value corresponding to its key.
	// ?? the latter one is whether the key exists
	val, exists := d[word]

	if !exists {
		return "", ErrNotFound
	}

	return val, nil
}

// ** you can MUTATE a map without passing an address to it like structs
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if exists := err == nil; exists {
		return ErrWordExists
	}

	if err == ErrNotFound {
		d[word] = definition
		return nil
	}

	return err
}
