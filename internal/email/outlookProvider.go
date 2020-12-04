package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/albb-b2b/push2b/pkg"
)

type outlookSender struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func (outlook outlookSender) Process(work pkg.Work) error {
	email := payload{}
	err := json.Unmarshal([]byte(work.Payload), &email)

	if err != nil {
		return err
	}

	if email.To == nil || len(email.To) <= 0 {
		return errors.New("to parameters is empty")
	}

	msg := formatRfc822(email.To[0], email.Body, email.Subject)
	return outlook.Send(msg, email.To, email.CC, email.Bcc)
}

func (outlook outlookSender) Send(body string, to []string, cc []string, bcc []string) error {

	auth := loginAuthfn("", outlook.username, outlook.password, outlook.host)

	e := smtp.SendMail(fmt.Sprintf("%s:%d", outlook.host, outlook.port), auth, outlook.username, to, []byte(body))
	if e != nil {
		fmt.Print(e)
	}
	return nil
}
