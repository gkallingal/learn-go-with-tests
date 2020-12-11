package maps

const (
	ErrNotFound     = DictionaryErr("word not found")
	ErrWordExists   = DictionaryErr("word exists")
	ErrCannotUpdate = DictionaryErr("cannot update: word not found")
	ErrCannotDelete = DictionaryErr("cannot delete: word not found")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrCannotUpdate
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
