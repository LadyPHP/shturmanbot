package main

import (
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"
	"os"
)

type Config struct {
	TelegramBotToken string
	DebugMode bool
}

// метод возвращающий состоятие приложения
// чтобы при прям запросе по https, приложение что-то возвращало, иначе на heroku падает
func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Bot is running."))
}

func setMsg() (string)  {
	msg := "Я пока умею отвечать только это."
	return msg
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = configuration.DebugMode

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updates, err := bot.GetUpdatesChan(u)
	updates := bot.ListenForWebhook("/" + bot.Token)

	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/", MainHandler)
	go http.ListenAndServe(":" + os.Getenv("PORT"), nil)

	// В канал updates будут приходить все новые сообщения.
	for update := range updates {
		// Создав структуру - можно её отправить обратно боту
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, setMsg())
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}