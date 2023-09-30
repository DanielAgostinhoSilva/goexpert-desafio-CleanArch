package events

type EventDispatcher interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(event Event) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear()
}
