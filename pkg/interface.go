package pkg

type Listener interface {
	//this method is to process the work before sending to job executer.
	//eg: consider an email sender trying to send email in standard format.
	//this method processes the format before executing the main job of email sender.
	//NOTE: (We have email senders that are message listener too!)
	Process(work Work) error
}

type Work struct {
	//used for transfering work headers:
	//eg: RabbitMQ routingkey or exchange can be transfered via this field
	Header string

	//used for transfering work payload.
	//eg: RabbitMQ json message can be transfered to internal system jobs via this field
	Payload string
}
