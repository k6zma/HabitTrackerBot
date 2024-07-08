package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandlingMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	switch inputMessage.Command() {
	case "start":
		startCommand(bot, inputMessage)
	default:
		fmt.Println("LOL")
	}
}
