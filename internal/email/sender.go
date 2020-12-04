//Package sender provides many types of implementation of various senders, like gmail-sender, outlook-sender, sms-sender
//There will be some standard formatters for email contents eg: formatrfc822
package sender

import "github.com/albb-b2b/push2b/pkg"

//many types of senders can be written in the internal package.
//currently there are two email-senders! (gmail, outlook)
//each of senders can implement their own send behaivour. (Sms Sender, Email Sender, Push Notification Sender etc...)

type Sender interface {
	//NOTE: each sender is a listener! because before sending a message it should listen to something to receive the prerequisites
	pkg.Listener
	Send(msg string, to []string, cc []string, bcc []string) error
}
