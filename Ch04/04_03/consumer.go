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

	c.Subscribe("events", func(m *nats.Msg) {
		var e Event
		if err := json.Unmarshal(m.Data, &e); err != nil {
			log.Printf("error: can't unmarshal - %s", err)
			return
		}
		log.Printf("got: %+v\n", e)
	})

	select {}
}
