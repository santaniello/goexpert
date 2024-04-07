package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispactcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispactcher {
	return &EventDispactcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispactcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errors.New("handler already registered")
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispactcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ed *EventDispactcher) Dispatch(event EventInterface) error {
	wg := &sync.WaitGroup{}
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}

func (ed *EventDispactcher) Remove(eventName string, handler EventHandlerInterface) error {
	if handlers, ok := ed.handlers[eventName]; ok {
		for i, h := range handlers {
			if h == handler {
				ed.handlers[eventName] = append(handlers[:i], handlers[i+1:]...)
				return nil
			}
		}
	}

	return errors.New("handler not found")

}

func (ed *EventDispactcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
