package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElementFromArray(t *testing.T) {
	t.Parallel()

	var testArr = make([]int, 10)
	for i := range 10 {
		testArr[i] = i
	}

	testArr = RemoveElementFromArray(testArr, 5)
	assert.Len(t, testArr, 9)
	for i := range 9 {
		assert.NotEqual(t, 5, testArr[i])
	}

	testArr = RemoveElementFromArray(testArr, 1)
	assert.Len(t, testArr, 8)
	for i := range 8 {
		assert.NotEqual(t, 1, testArr[i])
	}

	testArr = RemoveElementFromArray(testArr, 10)
	assert.Len(t, testArr, 8)
}
