package sender

import (
	"fmt"
	"net/smtp"

	"github.com/albb-b2b/push2b/pkg"
)

type googleSender struct {
	outlookSender
	identity string
}

type outlookSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewGoogleSender(host string, port int, username string, password string, identity string, from string) Sender {
	sender := googleSender{}
	sender.host = host
	sender.port = port
	sender.username = username
	sender.password = password
	sender.from = from

	return &sender
}

func (google googleSender) Process(work pkg.Work) error {
	to := make([]string, 0)
	to = append(to, "solhi.amir1371@gmail.com")

	return google.Send("subject", work.Payload, to, nil, nil)
}

func (google googleSender) Send(subject string, body string, to []string, cc []string, bcc []string) error {
	auth := smtp.PlainAuth(google.identity, google.username, google.password, google.host)

	e := smtp.SendMail(fmt.Sprintf("%s:%d", google.host, google.port), auth, google.username, to, []byte(body))
	if e != nil {
		fmt.Print(e)
	}
	return nil
}

func NewOutlookSender(host string, port int, username string, password string, from string) Sender {
	return &outlookSender{host, port, username, password, from}
}

func (outlook outlookSender) Process(work pkg.Work) error {
	to := make([]string, 0)
	to = append(to, "a.solhi@alibaba.ir")

	return outlook.Send("subject", work.Payload, to, nil, nil)
}

func (outlook outlookSender) Send(subject string, body string, to []string, cc []string, bcc []string) error {

	auth := loginAuthfn("", outlook.username, outlook.password, outlook.host)

	e := smtp.SendMail(fmt.Sprintf("%s:%d", outlook.host, outlook.port), auth, outlook.username, to, []byte(body))
	if e != nil {
		fmt.Print(e)
	}
	return nil
}
