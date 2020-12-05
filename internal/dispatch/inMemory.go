package dispatch

import (
	"log"
	"strings"

	"git.alibaba.ir/b2b/back/push2b/pkg"
)

//a simple dispather that routes incoming works to each listener stored on memory
type inMemoryDispatcher struct {
	listeners map[string]pkg.Listener
}

func NewInMemoryDispatcher() Dispatcher {
	dispatcher := inMemoryDispatcher{}
	dispatcher.listeners = make(map[string]pkg.Listener)

	return &dispatcher
}

//here we can change the work before sending it to be processed with listener.
func (d inMemoryDispatcher) Dispatch(listener pkg.Listener, work pkg.Work) {
	listener.Process(work)
}

func (d inMemoryDispatcher) AddListener(name string, listener pkg.Listener) {
	if d.listeners[name] != nil {
		panic("another listener added before with this name: " + name)
	}

	d.listeners[name] = listener
}

//dispatcher will be aware of incomming works in through this method.
//simple strategy is implemented to route the works and match available listeneres.
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
		log.Print("no listeners found to process job: ", work.Header)
	} else {
		d.Dispatch(found, work)
	}
}
