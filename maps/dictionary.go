package maps

import "errors"

var ErrNotFound = errors.New("word not found")

type Dictionary map[string]string

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return definition, ErrNotFound
	}
	return definition, nil
}
