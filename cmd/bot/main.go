package main

import (
	"log"
	"os"

	"github.com/Disaldr/go_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	for update := range updates {
		if update.Message != nil { // If we got a message
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehaivour(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I can't help you")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, productService *product.Service) {
	out := "All products: \n\n"
	products := productService.List()
	for _, p := range products {
		out += p.Title
		out += "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, out)
	bot.Send(msg)
}

func defaultBehaivour(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("\033[31m[%s] %s\033[0m", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	bot.Send(msg)
}
