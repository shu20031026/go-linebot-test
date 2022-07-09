package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
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

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.POST("/callback", func(ctx *gin.Context) {
		events, err := bot.ParseRequest(ctx.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		// 出席の場合
		var attendanceText string = "出席"

		var attendanceResponse string = "出席を確認しました"

		// 欠席の場合
		var absenceText string = "欠席"

		var absenceResponse string = "欠席を確認しました"

		//　出席と欠席の両方が含まれていた場合
		var bothResponse string = "出席または欠席のいずれかを送信してください"

		for _, event := range events {
			// イベントがメッセージの受信だった場合に実行
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {

				// メッセージがテキスト形式の場合に実行
				case *linebot.TextMessage:
					replyMessage := message.Text

					if strings.Contains(replyMessage, attendanceText) && strings.Contains(replyMessage, absenceText) {
						// メッセージに出席と欠席両方が含まれていた場合
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bothResponse)).Do()

					} else if strings.Contains(replyMessage, attendanceText) {
						// メッセージに出席が含まれていた場合
						//------------------ここでDBに出席と送信----------------------
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(attendanceResponse)).Do()

					} else if strings.Contains(replyMessage, absenceText) {
						// メッセージに欠席が含まれていた場合
						//------------------ここでDBに欠席と送信----------------------
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(absenceResponse)).Do()
					}
					// 上記以外は、おうむ返しで返信
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	router.Run(":" + port)
}
