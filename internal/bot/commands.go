package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func startCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Привет! Я бот для отслеживания привычек. "+
		"Я помогу вам добавить новые привычки, отмечать их выполнение и следить за прогрессом.\n"+
		"Воспользуйтесь командой /help, чтобы посмотреть все доступные команды.",
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
