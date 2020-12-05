package sender

import (
	"fmt"

	"git.alibaba.ir/b2b/back/push2b/pkg"
)

//This struct is used to deserialize the payload of works routed to special sender.
type payload struct {
	To      []string `json:"to"`
	CC      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	Body    string   `json:"body"`
	Subject string   `json:"subject"`
}

//instanciates an email sender based on this Outlook smtp server.
func NewOutlookSender(config pkg.OutlookConfig) Sender {
	return &outlookSender{config.Host, config.Port, config.Username, config.Password, config.Username}
}

//instanciates a sender based on the Gmail smtp server.
func NewGoogleSender(config pkg.GmailConfig) Sender {
	sender := googleSender{}
	sender.host = config.Host
	sender.port = config.Port
	sender.username = config.Username
	sender.password = config.Password
	sender.from = config.Username

	return &sender
}

//The msg parameter for sending email should be an RFC 822-style email with headers-first,
//a blank line, and then the message body. The lines of msg should be CRLF terminated.
//The msg headers should usuallyinclude fields such as "From", "To", "Subject", and "Cc".
//Sending "Bcc" messages is accomplished by including an email address in the to
//parameter but not including it in the msg headers.
func formatRfc822(to string, body string, subject string) string {
	return fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body)
}
