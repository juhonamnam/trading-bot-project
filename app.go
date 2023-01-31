package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/juhonamnam/telego"
	"github.com/juhonamnam/trading-bot-project/tradingalgorithm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Cannot read .env file")
		panic(err.Error())
	}
	bot := telego.Initialize(os.Getenv("TELEGRAM_BOT_API_KEY"))
	tradingalgorithm.SetMessengerBot(bot, os.Getenv("VBS_CHAT_ID"))
	tradingalgorithm.Start()
	// bot.SetUpdateHandler(func(ctx *telego.Context) {})
	// bot.Start()
}
