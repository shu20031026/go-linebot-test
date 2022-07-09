package model

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func LineBotInit() *linebot.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	lineBotChannelSecret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	lineBotChannelToken := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	bot, err := linebot.New(
		lineBotChannelSecret,
		lineBotChannelToken,
	)

	if err != nil {
		log.Fatal(err)
	}
	return bot
}
