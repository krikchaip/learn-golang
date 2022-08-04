package maps

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// ?? map-lookup returns 2 values
	val, exists := d[word]

	if !exists {
		return "", errors.New("could not find the word you were looking for")
	}

	return val, nil
}
