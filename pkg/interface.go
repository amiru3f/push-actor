package pkg

type Listener interface {
	Process(work Work) error
}

type Work struct {
	Header  string
	Payload string
}
