package maps

type Dictionary map[string]string

func Search(m Dictionary, key string) string {
	return m[key]
}
