package main

import (
	// "go-telegram-bot-examples/echo_bot"
	"go-telegram-bot-examples/echo_bot_gin"
	"os"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	webhookBase := os.Getenv("WEBHOOK")
	// echo_bot.Run(token, webhookBase, "0.0.0.0", "80", true)
	echo_bot_gin.Run(token, webhookBase, "0.0.0.0", "8080", false)
}
