package events

import "sync"

type EventHandler interface {
	Handle(event Event, wg *sync.WaitGroup)
}
