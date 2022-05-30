package main

import (
	"go-telegram-bot-examples/echo_bot"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	webhookBase := "https://adaa-8-218-73-235.ap.ngrok.io/"
	echo_bot.Run(token, webhookBase, "0.0.0.0:80", true)
}
