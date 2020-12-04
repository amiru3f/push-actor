//Package consume is related to all of consuming behaviours.
//for instance it consumes rabbit, receives jobs and uses dispatcher to route.
package consume

import (
	"fmt"
	"log"
)

type Consumer interface {
	Consume() error
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func consumingInterrupted() {
	fmt.Print("connection/channel closed")
}
