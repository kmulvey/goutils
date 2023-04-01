package goutils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randomSource *rand.Rand

func init() {
	//nolint:gosec
	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[randomSource.Intn(len(letters))]
	}
	return string(s)
}
