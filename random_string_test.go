package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Parallel()

	var a = RandomString(10)
	assert.Equal(t, 10, len(a))

	var b = RandomString(10)
	assert.Equal(t, 10, len(a))

	assert.NotEqual(t, a, b)
}
