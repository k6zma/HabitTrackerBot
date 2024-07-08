package utils

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	resetColor   = "\033[0m"
	redColor     = "\033[31m"
	greenColor   = "\033[32m"
	yellowColor  = "\033[33m"
	blueColor    = "\033[34m"
	magentaColor = "\033[35m"
	cyanColor    = "\033[36m"
	grayColor    = "\033[37m"
	whiteColor   = "\033[97m"
	boldRed      = "\033[31;1m"
	boldGreen    = "\033[32;1m"
	boldYellow   = "\033[33;1m"
	boldBlue     = "\033[34;1m"
	boldMagenta  = "\033[35;1m"
	boldCyan     = "\033[36;1m"
	boldGray     = "\033[37;1m"
	boldWhite    = "\033[97;1m"
)

func PanicWithColor(msg interface{}) {
	panic(fmt.Sprintf("%s%s%s\n", redColor, msg, resetColor))
}

func SuccessfulExecution(msg string) {
	fmt.Printf("%sSuccessful execution function:%s %s%s%s\n", yellowColor, resetColor, greenColor, msg, resetColor)
}

func SuccessfulAuthorizing(bot *tgbotapi.BotAPI) {
	fmt.Printf("%sAuthorized on bot account:%s %s%s%s\n", cyanColor, resetColor, boldBlue, bot.Self.UserName, resetColor)
}
