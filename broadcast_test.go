package goutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnsubscribe(t *testing.T) {
	var broadcast = NewBroadcast()
	var subscriberOne = broadcast.Subscribe()
	_ = broadcast.Subscribe()
	broadcast.Unsubscribe(subscriberOne)
	assert.Equal(t, 1, len(broadcast.Subscribers))
}

func TestBroadcast(t *testing.T) {
	var broadcast = NewBroadcast()
	var subscriberOne = broadcast.Subscribe()
	var subscriberTwo = broadcast.Subscribe()
	var done = make(chan struct{})
	var message = "test message"

	go func() {
		var oneOpen, twoOpen = true, true
		var numMessages uint8
		for oneOpen || twoOpen {
			select {
			case msg, open := <-subscriberOne:
				if !open {
					oneOpen = false
					continue
				}
				assert.Equal(t, message, msg)
				numMessages++
			case msg, open := <-subscriberTwo:
				if !open {
					twoOpen = false
					continue
				}
				assert.Equal(t, message, msg)
				numMessages++
			}
		}
		assert.Equal(t, uint8(2), numMessages)
		close(done)
	}()

	broadcast.Broadcast(message)
	broadcast.Shutdown()
	<-done
}
