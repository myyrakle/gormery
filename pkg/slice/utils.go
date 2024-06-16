package slice

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s Strings) Contains(e string) bool {
	return Contains[string](s, e)
}
