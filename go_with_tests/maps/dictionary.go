package maps

import "errors"

var errorNotFound = errors.New("Could not find the word you are looking for")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errorNotFound
	}
	return definition, nil
}
