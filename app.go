package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	err := godotenv.Load("bot.env")
	if err != nil {
		panic(err)
	}
	r := gin.New()
	r.POST("/callback", callbackHandler)
	http.Handle("/", r)
}

func callbackHandler(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"), os.Getenv("LINE_BOT_CHANNEL_TOKEN"), linebot.WithHTTPClient(urlfetch.Client(ctx)))
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		log.Errorf(ctx, "Error: %v", err)
	}
	for _, event := range events {
		replyToken := event.ReplyToken
		if event.Type != linebot.EventTypeMessage {
			_, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage("キャラクターの名前を送信してね♪")).Do()
			if err != nil {
				log.Errorf(ctx, "Error: %v", err)
			}
			continue
		}
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			_, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.Text)).Do()
			if err != nil {
				log.Errorf(ctx, "Error: %v", err)
			}
		default:
			_, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage("キャラクターの名前を送信してね♪")).Do()
			if err != nil {
				log.Errorf(ctx, "Error: %v", err)
			}
		}
	}
}
