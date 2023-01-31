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
var _chatId string

func SetMessengerBot(bot messengerBot, chatId string) {
	_messengerBot = bot
	_chatId = chatId
}

func sendMessage(text string) {
	_messengerBot.Request("sendMessage", message{
		ChatId:    _chatId,
		Text:      text,
		ParseMode: "HTML",
	})
}
