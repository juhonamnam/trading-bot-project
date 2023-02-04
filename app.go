package main

import (
	"github.com/juhonamnam/telego"
	"github.com/juhonamnam/trading-bot-project/env"
	"github.com/juhonamnam/trading-bot-project/logger"
	"github.com/juhonamnam/trading-bot-project/tradingalgorithm"
)

func main() {
	bot := telego.Initialize(env.BotAPIKey)
	bot.SetLogger(logger.GetTelegoLogger())
	tradingalgorithm.SetMessengerBot(bot)
	tradingalgorithm.Start()
	// bot.SetUpdateHandler(func(ctx *telego.Context) {})
	// bot.Start()
}
