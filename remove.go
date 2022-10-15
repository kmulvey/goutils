package goutils

// RemoveElementFromArray does just that, if an index is passed in
// that is larger than the slice length then the original slice is returned.
func RemoveElementFromArray[T any](slice []T, s int) []T {
	if s >= len(slice) {
		return slice
	}
	return append(slice[:s], slice[s+1:]...)
}
