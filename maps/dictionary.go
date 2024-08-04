package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	var found, ok = d[word]

	if !ok {
		return found, ErrNotFound
	}

	return found, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
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
