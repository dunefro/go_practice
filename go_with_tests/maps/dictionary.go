package maps

import "errors"

var (
	errorNotFound   = errors.New("Could not find the word you are looking for")
	errorWordExists = errors.New("key already exists in the map")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	if err == nil {
		return errorWordExists
	}
	d[key] = value
	return nil
}
