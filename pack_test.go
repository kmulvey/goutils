package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack(t *testing.T) {
	t.Parallel()

	var store = Pack(1111, 2222)
	var valA, valB = Unpack(store)
	assert.Equal(t, int64(1111), valA)
	assert.Equal(t, int64(2222), valB)
}
