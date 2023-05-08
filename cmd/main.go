package main

import (
	"log"
	"os"
	tgClient "password-storage-bot/internal/app/clients/telegram"
	event_consumer "password-storage-bot/internal/app/consumer/event-consumer"
	"password-storage-bot/internal/app/events/telegram"
	"password-storage-bot/internal/app/storage/postgres"
)

const batchSize = 100

func main() {
	s, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(os.Getenv("TG_BOT_HOST"), os.Getenv("TG_BOT_TOKEN")),
		s,
	)

	log.Print("bot started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
