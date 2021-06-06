package maps

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}
	d[key] = value
	return nil
}

// Implementing the error interface
func (e DictionaryErr) Error() string {
	return string(e)
}
