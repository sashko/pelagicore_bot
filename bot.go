package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func replyHelp() string {
	str := "I accept the following commands:\n\n" +
		"/help\t    print available commands"

	return str
}

func main() {
	log.Println("pelagicore_bot has started")

	bot, err := tgbotapi.NewBotAPI("")
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
			msgText = replyHelp()
		default:
			msgText = "Implement this command first mate :/"
		}

	send_msg:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		bot.Send(msg)
	}
}
