package tradingalgorithm

type messengerBot interface {
	Request(endpoint string, data any) (*[]byte, error)
}

type message struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

var _messengerBot messengerBot = nil

func SetMessengerBot(bot messengerBot) {
	_messengerBot = bot
}

func sendMessage(text string, chatId string) {
	_messengerBot.Request("sendMessage", message{
		ChatId:    chatId,
		Text:      text,
		ParseMode: "HTML",
	})
}
