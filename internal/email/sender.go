package sender

import "github.com/albb-b2b/push2b/pkg"

type Sender interface {
	pkg.Listener
	Send(subject string, body string, to []string, cc []string, bcc []string) error
}
