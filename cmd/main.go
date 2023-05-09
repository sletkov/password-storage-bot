package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	tgClient "password-storage-bot/internal/app/clients/telegram"
	event_consumer "password-storage-bot/internal/app/consumer/event-consumer"
	"password-storage-bot/internal/app/events/telegram"
	"password-storage-bot/internal/app/storage/postgres"
)

const batchSize = 100

func init() {
	// Loads variables from .env
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	s, err := postgres.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(os.Getenv("TG_BOT_HOST"), os.Getenv("TG_BOT_TOKEN")),
		s,
	)

	//// Delete messages logic
	//bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	//if err != nil {
	//	log.Fatal("can't create botApi: ", err)
	//}
	//
	//delegeMsg := tgbotapi.DeleteMessageConfig{
	//	ChatID:    chatID,
	//	MessageID: messageId,
	//}
	//
	//_, err = bot.DeleteMessage(deleteMsg)
	//if err != nil {
	//	log.Println(err)
	//}

	log.Print("bot started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
