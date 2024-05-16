package goutils

import (
	"math/rand"
	"time"
)

// RandomString returns a cases sensitive random string of the given length. The randonSource and
// character array are in the function to avoid data races. At this time its overkill to store them
// in a type but this can be done in the future if necessary.
func RandomString(n int) string {
	//nolint:gosec
	var randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, n)
	for i := range s {
		s[i] = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")[randomSource.Intn(52)]
	}
	return string(s)
}
