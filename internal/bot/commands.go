package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func startCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Привет! Я бот для отслеживания привычек. "+
		"Я помогу вам добавить новые привычки, отмечать их выполнение и следить за прогрессом.\n\n"+
		"Воспользуйтесь командой /help, чтобы посмотреть все доступные команды.",
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Привет! Я бот для отслеживания привычек. Вот список всех доступных команд, которые помогут вам управлять своими привычками:\n\n"+
		"- /start - Приветствие и объяснение функционала бота.\n"+
		"- /help - Список всех доступных команд и их описание.\n"+
		"- /add_habit <название> - Добавить новую привычку. Например: /add_habit чтение\n"+
		"- /delete_habit <название> - Удалить привычку. Например: /delete_habit чтение\n"+
		"- /list_habits - Показать список всех добавленных привычек.\n"+
		"- /mark_habit <название> - Отметить выполнение привычки. Например: /mark_habit чтение\n"+
		"- /unmark_habit <название> - Снять отметку выполнения привычки. Например: /unmark_habit чтение\n"+
		"- /stats - Показать статистику выполнения всех привычек (количество выполнений).\n\n"+
		"Используйте эти команды, чтобы отслеживать и управлять своими привычками. Начните с добавления первой привычки с помощью команды /add_habit <название>!",
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
