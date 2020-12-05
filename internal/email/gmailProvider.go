package sender

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/smtp"

	"git.alibaba.ir/b2b/back/push2b/pkg"
)

type googleSender struct {
	outlookSender
	identity string
}

func (google googleSender) Process(work pkg.Work) error {
	email := payload{}
	err := json.Unmarshal([]byte(work.Payload), &email)

	if err != nil {
		return err
	}

	if email.To == nil || len(email.To) <= 0 {
		return errors.New("to parameters is empty")
	}

	msg := formatRfc822(email.To[0], email.Body, email.Subject)
	return google.Send(msg, email.To, email.CC, email.Bcc)
}

func (google googleSender) Send(body string, to []string, cc []string, bcc []string) error {
	auth := smtp.PlainAuth(google.identity, google.username, google.password, google.host)

	e := smtp.SendMail(fmt.Sprintf("%s:%d", google.host, google.port), auth, google.username, to, []byte(body))
	if e != nil {
		fmt.Print(e)
	}
	return nil
}
