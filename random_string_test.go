package goutils

import (
	"sync"
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

// make sure there is no data race
func TestRandomStringParallel(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < 1000; i++ {
				var a = RandomString(10)
				assert.Equal(t, 10, len(a))
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
