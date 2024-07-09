package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	habit "github.com/k6zma/HabitTrackerBot/internal/model"
)

func startCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Привет! Я бот для отслеживания привычек. "+
		"Я помогу вам добавить новые привычки, отмечать их выполнение и следить за прогрессом.\n\n"+
		"Воспользуйтесь командой /help, чтобы посмотреть все доступные команды.",
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Привет! Я бот для отслеживания привычек. Вот список всех доступных команд, которые помогут вам управлять своими привычками:\n\n"+
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

func addHabitCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	habitName := inputMessage.CommandArguments()
	if habitName == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Пожалуйста, укажите название привычки.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	if habit.Users == nil {
		habit.Users = make(map[int64]*habit.User)
	}

	user, exists := habit.Users[inputMessage.Chat.ID]
	if !exists {
		user = &habit.User{Habits: make(map[string]*habit.Habit)}
		habit.Users[inputMessage.Chat.ID] = user
	}

	if _, exists := user.Habits[habitName]; exists {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Эта привычка уже добавлена.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	user.Habits[habitName] = &habit.Habit{Name: habitName}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Привычка: %v добавлена.", habitName),
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func deleteHabitCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	habitName := inputMessage.CommandArguments()
	if habitName == "" {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Пожалуйста, укажите название привычки.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	if habit.Users == nil {
		habit.Users = make(map[int64]*habit.User)
	}

	user, exists := habit.Users[inputMessage.Chat.ID]
	if !exists {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Вы еще не добвыили ни одной привычки.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	if _, exists := user.Habits[habitName]; !exists {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Этой привычки не существует.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	delete(user.Habits, habitName)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Привычка: %v - удалена.", habitName),
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func listHabitsCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	if habit.Users == nil {
		habit.Users = make(map[int64]*habit.User)
	}

	user, exists := habit.Users[inputMessage.Chat.ID]
	if !exists {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Вы еще не добавили ни одной привычки.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	if len(user.Habits) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "У вас нет добавленных привычек.")
		msg.ReplyToMessageID = inputMessage.MessageID

		bot.Send(msg)
		return
	}

	habitsList := "Список ваших привычек:\n"
	for _, habit := range user.Habits {
		habitsList += fmt.Sprintf("- %s\n", habit.Name)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, habitsList)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func defaultCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Введена неправильная команда: %v", inputMessage.Text),
	)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
