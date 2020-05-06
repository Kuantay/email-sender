package owner

import (
	"encoding/json"
	"log"

	config "github.com/AkezhanOb1/email-sender/configs"
	"github.com/AkezhanOb1/email-sender/model"
	"github.com/AkezhanOb1/email-sender/services/email"
	"github.com/streadway/amqp"
)

//SendEmailToBusinessOwner is a func
func SendEmailToBusinessOwner() {
	conn, err := amqp.Dial(config.RabbitConnection)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"businessRegistration", // name
		false,                  // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		log.Println(err)
	}

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Println(err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msg {
			var e model.EmailSender
			err = json.Unmarshal(d.Body, &e)
			if err != nil {
				log.Println(err)
			}

			err = email.WelcomeEmail(e.Email.Address)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
