package tradingalgorithm

import (
	"encoding/json"

	"github.com/juhonamnam/trading-bot-project/logger"
)

type messengerBot interface {
	Request(endpoint string, data any) (*[]byte, error)
}

type message struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type baseResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

var _messengerBot messengerBot = nil

func SetMessengerBot(bot messengerBot) {
	_messengerBot = bot
}

func sendMessage(text string, chatId string) {
	res, err := _messengerBot.Request("sendMessage", message{
		ChatId:    chatId,
		Text:      text,
		ParseMode: "HTML",
	})
	if err != nil {
		logger.Telego.Error(err.Error())
	}

	var result baseResponse
	err = json.Unmarshal(*res, &result)
	if err != nil {
		logger.Telego.Error(err.Error())
	}

	if !result.Ok {
		logger.Telego.Error(result.Description)
	}
}
