package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

var startMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("trip"),
		tgbotapi.NewKeyboardButton("help"),
	),
)
var roleMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("driver"),
		tgbotapi.NewKeyboardButton("passenger"),
	),
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

	states := make(map[string]string)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = debug

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	for update := range bot.GetUpdatesChan(updateConfig) {
		username := update.Message.From.UserName
		fmt.Println(username)

		state, exists := states[username]
		fmt.Println(state)

		if !exists {
			states[username] = "start"
		}

		message := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		text := update.Message.Text

		switch states[username] {
		case "start":
			switch text {
			case "trip":
				states[username] = "trip"
				message.Text = "Are you a driver or a passenger?"
				message.ReplyMarkup = roleMenu
			default:
				message.Text = "What do you want to do?"
				message.ReplyMarkup = startMenu
			}
		case "trip":
			states[username] = "role"

			switch text {
			case "driver":
				message.Text = "How many people do you want to take?"
				message.ReplyMarkup = numericKeyboard
			case "passenger":
				message.Text = "How many people will go?"
				message.ReplyMarkup = numericKeyboard
			}
		case "role":
			message.Text = fmt.Sprintf("%s: %s", states[username], text)
			states[username] = "start" // just for now
		}

		if _, err := bot.Send(message); err != nil {
			fmt.Println(err)
		}
	}
}