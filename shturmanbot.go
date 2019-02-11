package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	TelegramBotToken string
	DebugMode        bool
}

// метод возвращающий состоятие приложения
// чтобы при прям запросе по https, приложение что-то возвращало, иначе на heroku падает
func MainHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Bot is running."))
}

func GetBusSchedule(from, to string) (link string) {
	when := time.Now()
	link = fmt.Sprintf("https://t.rasp.yandex.ru/search/bus/?fromName=%s&toName=%s&when=%s", from, to, when)

	return link
}

func GetBalance() (msg string, err error) {
	// TODO: убрать этот хардкод
	api := "http://strelkacard.ru/api/cards/status/?cardnum=03330921822&cardtypeid=3ae427a1-0f17-4524-acb1-a3f50090a8f3"

	res, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", response)

	var dat map[string]interface{}
	if err := json.Unmarshal(response, &dat); err != nil {
		panic(err)
	}

	num := dat["balance"].(float64) / 100
	msg = fmt.Sprintf("Баланс в порядке %f", num)
	if num < 200 {
		msg = fmt.Sprintln("Пополнить баланс?")
		// TODO: добавить команду пополнения баланса
	}

	return msg, nil
}

func setMsg() string {
	rasp := GetBusSchedule("Москва", "Звенигород")
	balance, _ := GetBalance()
	msg := fmt.Sprintf("Расписание %s . Баланс %s", rasp, balance)
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
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	// В канал updates будут приходить все новые сообщения.
	for update := range updates {
		// Создав структуру - можно её отправить обратно боту
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, setMsg())
		bot.Send(msg)
	}
}
