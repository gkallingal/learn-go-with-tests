package maps

import "errors"

var (
	ErrNotFound   = errors.New("word not found")
	ErrWordExists = errors.New("word exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return definition, ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[word] = definition
	default:
		return err
	}
	return nil
}
