package maps

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if exists := err == nil; exists {
		d[word] = definition
		return nil
	}

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

type DictionaryErr string

// ** error type can be anything as long as Error() is implemented
func (e DictionaryErr) Error() string {
	return string(e)
}

// ?? error constants pattern
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)
