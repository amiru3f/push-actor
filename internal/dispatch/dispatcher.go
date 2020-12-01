package dispatch

import "github.com/albb-b2b/push2b/pkg"

type Dispatcher interface {
	Received(work pkg.Work)
	Dispatch(listener pkg.Listener, work pkg.Work)
	AddListener(name string, listener pkg.Listener)
}
