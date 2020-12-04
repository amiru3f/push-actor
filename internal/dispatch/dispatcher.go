//all of system incomming work/jobs are routed here to avaialble executers.
package dispatch

import "github.com/albb-b2b/push2b/pkg"

//may have many implementations based on environment.
//eg: InMemoryDispatcher
//eg: ReflectiveDispatcher
//eg: RedisBasedDispatcher (used like inmemory dispatcher but it uses redis to save listener types)
type Dispatcher interface {
	//to be awared of incomming job stream.
	Received(work pkg.Work)
	//to dispatch a work/job to specific listener
	Dispatch(listener pkg.Listener, work pkg.Work)
	//registers new listeners with specific name and type.
	AddListener(name string, listener pkg.Listener)
}
