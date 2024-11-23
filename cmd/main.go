package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println(err)
	}
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = debug

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	for update := range bot.GetUpdatesChan(updateConfig) {
		message := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if _, err := bot.Send(message); err != nil {
			fmt.Println(err)
		}
	}
}