package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeChannels(t *testing.T) {
	t.Parallel()

	var chans = []chan int{make(chan int, 5), make(chan int, 5)}
	chans[0] <- 0
	chans[0] <- 1
	chans[0] <- 2
	chans[0] <- 3
	chans[0] <- 4
	chans[1] <- 5
	chans[1] <- 6
	chans[1] <- 7
	chans[1] <- 8
	chans[1] <- 9
	close(chans[0])
	close(chans[1])

	var expected = map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}}
	for val := range MergeChannels(chans...) {
		delete(expected, val)
	}
	assert.Empty(t, expected)
}
