package main

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURL := os.Getenv("NATS_URL")
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subjects := []string{
		"product.user.set",
		"product.user.get",
		"product.user.update",
		"product.user.delete",
	}

	messages := []string{
		"Set user data",
		"Get user data",
		"Update user data",
		"Delete user data",
	}

	for i, subject := range subjects {
		if subject == "product.user.get" || subject == "product.user.update" {
			msg, err := nc.Request(subject, []byte(messages[i]), nats.DefaultTimeout)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Received reply on subject %s: %s", subject, string(msg.Data))
		} else {
			err = nc.Publish(subject, []byte(messages[i]))
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Published message on subject %s: %s", subject, messages[i])
		}
	}
}
