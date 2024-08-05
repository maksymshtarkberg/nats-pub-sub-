package main

import (
	"fmt"
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

	for _, subject := range subjects {
		_, err := nc.Subscribe(subject, func(msg *nats.Msg) {
			log.Printf("Received a message on subject %s: %s", msg.Subject, string(msg.Data))
			if msg.Reply != "" {
				replyMsg := fmt.Sprintf("Reply to %s", msg.Subject)
				msg.Respond([]byte(replyMsg))
			}
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Subscriber is listening...")
	select {}
}
