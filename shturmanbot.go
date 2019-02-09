package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(configuration.TelegramBotToken)

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	if err != nil {
		log.Panic(err)
	}

	status_message, err := bot.GetMe()
	fmt.Println(status_message)
}