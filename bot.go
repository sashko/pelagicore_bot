package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

type Config struct {
	Token string `json:"token"`
}

func getConfig() Config {
	var config Config

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {
	log.Println("pelagicore_bot has started")

	var config Config

	config = getConfig()

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorization as %s SUCCESS", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.FirstName, update.Message.Text)

		msgText := ""

		if !update.Message.IsCommand() {
			msgText = "I am not a chat bot ya stupid!"
			goto send_msg
		}

		switch update.Message.Text {
		case "/help":
			msgText = handleHelp()
		default:
			msgText = "Implement this command first mate :/"
		}

	send_msg:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		bot.Send(msg)
	}
}
