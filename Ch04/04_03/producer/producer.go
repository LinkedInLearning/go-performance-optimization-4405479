package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Event struct {
	Time   time.Time `json:"time"`
	Login  string    `json:"login"`
	Action string    `json:"action"`
}

func main() {
	c, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		e := Event{
			Time:   time.Now(),
			Login:  "elliot",
			Action: "Login",
		}
		data, err := json.Marshal(e)
		if err != nil {
			log.Printf("error: can't encode - %s", err)
			continue
		}

		if err := c.Publish("events", data); err != nil {
			log.Fatalf("error: %s", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
