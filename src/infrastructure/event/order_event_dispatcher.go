package event

import (
	"errors"
	"github.com/DanielAgostinhoSilva/goexpert-desafio-CleanArch/src/domain/events"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type OrderEventDispatcher struct {
	handlers map[string][]events.EventHandler
}

func NewEventDispatcher() *OrderEventDispatcher {
	return &OrderEventDispatcher{
		handlers: make(map[string][]events.EventHandler),
	}
}

func (ev *OrderEventDispatcher) Dispatch(event events.Event) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *OrderEventDispatcher) Register(eventName string, handler events.EventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *OrderEventDispatcher) Has(eventName string, handler events.EventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *OrderEventDispatcher) Remove(eventName string, handler events.EventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func (ed *OrderEventDispatcher) Clear() {
	ed.handlers = make(map[string][]events.EventHandler)
}
