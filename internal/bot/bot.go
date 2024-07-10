package bot

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	habit "github.com/k6zma/HabitTrackerBot/internal/model"
	"github.com/k6zma/HabitTrackerBot/internal/utils"
)

func Start() {
	err := godotenv.Load(".env")
	if err != nil {
		utils.PanicWithColor(fmt.Sprintf("Error occrured while loading .env file: %v", err))
	}

	utils.SuccessfulExecution("reading .env file")

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		utils.PanicWithColor(err)
	}

	utils.SuccessfulExecution("connecting a bot using a token")

	bot.Debug = true

	utils.SuccessfulAuthorizing(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	if habit.Users == nil {
		habit.Users = make(map[int64]*habit.User)
	}

	for update := range updates {
		if update.Message != nil {
			HandlingMessage(bot, update.Message)
		}
	}
}
