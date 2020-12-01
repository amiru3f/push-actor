package main

import (
	"fmt"

	"github.com/albb-b2b/push2b/internal/consume"
	"github.com/albb-b2b/push2b/internal/dispatch"
	sender "github.com/albb-b2b/push2b/internal/email"
	"github.com/albb-b2b/push2b/pkg"
)

func main() {

	dispatcher := dispatch.NewInMemoryDispatcher()

	dispatcher.AddListener(pkg.MAIL_GMAIL, sender.NewGoogleSender("smtp.gmail.com", 587, "solhi.amir1371@gmail.com", "tvtaeifhwbtddsag", "", "solhi.amir1371@gmail.com"))
	dispatcher.AddListener(pkg.MAIL_OUTLOOK, sender.NewOutlookSender("smtp-mail.outlook.com", 587, "test_smtp22@outlook.com", "Mohsen0000", "test_smtp22@outlook.com"))

	consumer := consume.NewRabbitConsumer("localhost", 5672, "user", "bitnami", dispatcher)

	fmt.Errorf("error occured: %s", consumer.Consume())
}
