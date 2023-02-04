package env

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var BotAPIKey string
var VBSChatId string
var IsProduction bool

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Cannot read .env file")
		panic(err.Error())
	}
	BotAPIKey = os.Getenv("TELEGRAM_BOT_API_KEY")
	VBSChatId = os.Getenv("VBS_CHAT_ID")
	isProduction := flag.Bool("production", false, "Production Mode")
	flag.Parse()
	IsProduction = *isProduction
}
