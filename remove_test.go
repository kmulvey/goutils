package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElementFromArray(t *testing.T) {
	t.Parallel()

	var testArr = make([]int, 10)
	for i := 0; i < 10; i++ {
		testArr[i] = i
	}

	testArr = RemoveElementFromArray(testArr, 5)
	assert.Equal(t, 9, len(testArr))
	for i := 0; i < 9; i++ {
		assert.False(t, testArr[i] == 5)
	}

	testArr = RemoveElementFromArray(testArr, 1)
	assert.Equal(t, 8, len(testArr))
	for i := 0; i < 8; i++ {
		assert.False(t, testArr[i] == 1)
	}

	testArr = RemoveElementFromArray(testArr, 10)
	assert.Equal(t, 8, len(testArr))
}
