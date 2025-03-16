package goutils

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	t.Parallel()

	var a = RandomString(10)
	assert.Len(t, a, 10)

	var b = RandomString(10)
	assert.Len(t, b, 10)

	fmt.Printf("a: %s\nb: %s\n", a, b)
	assert.NotEqual(t, a, b)
}

// make sure there is no data race.
func TestRandomStringParallel(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for range 1000 {
				var a = RandomString(10)
				assert.Len(t, a, 10)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
