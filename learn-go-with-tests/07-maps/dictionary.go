package maps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// ?? map-lookup returns 2 values
	val, exists := d[word]

	if !exists {
		return "", ErrNotFound
	}

	return val, nil
}
