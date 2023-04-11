package slices

// Contains returns true if v is in s.
func Contains[T comparable](s []T, v T) bool {
	for _, vs := range s {
		if vs == v {
			return true
		}
	}

	return false
}
