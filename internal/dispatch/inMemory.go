package dispatch

import (
	"strings"

	"github.com/albb-b2b/push2b/pkg"
)

type inMemoryDispatcher struct {
	listeners map[string]pkg.Listener
}

func NewInMemoryDispatcher() Dispatcher {
	dispatcher := inMemoryDispatcher{}
	dispatcher.listeners = make(map[string]pkg.Listener)

	return &dispatcher
}

func (d inMemoryDispatcher) Dispatch(listener pkg.Listener, work pkg.Work) {
	listener.Process(work)
}

func (d inMemoryDispatcher) AddListener(name string, listener pkg.Listener) {
	if d.listeners[name] != nil {
		panic("another listener added before with this name: " + name)
	}

	d.listeners[name] = listener
}

func (d inMemoryDispatcher) Received(work pkg.Work) {

	var found pkg.Listener
	for k, l := range d.listeners {
		if strings.Contains(k, "gmail") && strings.Contains(work.Header, "gmail") {
			found = l
			break
		}

		if strings.Contains(k, "outlook") && strings.Contains(work.Header, "outlook") {
			found = l
			break
		}
	}

	if nil == found {
		panic("no listeners found to process job")
	}

	d.Dispatch(found, work)
}
