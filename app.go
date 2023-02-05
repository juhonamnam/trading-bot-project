package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/juhonamnam/telego"
	"github.com/juhonamnam/trading-bot-project/env"
	"github.com/juhonamnam/trading-bot-project/logger"
	"github.com/juhonamnam/trading-bot-project/tradingalgorithm"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		tradingalgorithm.Stop()
	}()

	bot := telego.Initialize(env.BotAPIKey)
	bot.SetLogger(logger.GetTelegoLogger())
	tradingalgorithm.SetMessengerBot(bot)
	tradingalgorithm.Start()
}
