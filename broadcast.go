package goutils

import "sync"

type Broadcast struct {
	Subscribers []chan any
	Lock        sync.RWMutex
}

func NewBroadcast() *Broadcast {
	return new(Broadcast)
}

func (b *Broadcast) Subscribe() chan any {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	var subscriber = make(chan any, 1000) // buffered so slow consumers dont block the others
	b.Subscribers = append(b.Subscribers, subscriber)
	return subscriber
}

func (b *Broadcast) Broadcast(message any) {
	b.Lock.RLock()
	defer b.Lock.RUnlock()

	for _, subscriber := range b.Subscribers {
		subscriber <- message
	}
}

func (b *Broadcast) Unsubscribe(unsubscriber chan any) {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	for i, subscriber := range b.Subscribers {
		if unsubscriber == subscriber {
			b.Subscribers = append(b.Subscribers[:i], b.Subscribers[i+1:]...)
			return
		}
	}
}

func (b *Broadcast) Shutdown() {
	b.Lock.Lock()
	defer b.Lock.Unlock()

	for _, subscriber := range b.Subscribers {
		close(subscriber)
	}
}
