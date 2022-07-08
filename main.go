package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func LordLineEnv() (string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	lineBotChannelSecret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	lineBotChannelToken := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	return lineBotChannelSecret, lineBotChannelToken
}

func main() {
	lineBotChannelSecret, lineBotChannelToken := LordLineEnv()

	bot, err := linebot.New(
		lineBotChannelSecret,
		lineBotChannelToken,
	)
	if err != nil {
		log.Fatal(err)
	}

	result := "test"

	message := linebot.NewTextMessage(result)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
