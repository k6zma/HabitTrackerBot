package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandlingMessage(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	switch inputMessage.Command() {
	case "start":
		startCommand(bot, inputMessage)
	case "help":
		helpCommand(bot, inputMessage)
	case "add_habit":
		addHabitCommand(bot, inputMessage)
	case "delete_habit":
		deleteHabitCommand(bot, inputMessage)
	case "list_habits":
		listHabitsCommand(bot, inputMessage)
	case "mark_habit":
		markHabbitCommand(bot, inputMessage)
	case "unmark_habit":
		unmarkHabitCommand(bot, inputMessage)
	default:
		defaultCommand(bot, inputMessage)
	}
}
