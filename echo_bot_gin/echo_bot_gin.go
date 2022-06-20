package echo_bot_gin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

var bot *tgbotapi.BotAPI

func initTelegram(token string, webhookBase string, debug bool) {

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhook(webhookBase + bot.Token)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
}

func webhookHandler(c *gin.Context) {
	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}
	//c.JSON(http.StatusOK, gin.H{})
	log.Printf("ChatID: %+v MessageID: %+v From: %+v Text: %+v \n", update.Message.Chat.ID, update.Message.MessageID, update.Message.From, update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	if update.Message.Text != "" {
		_, err = bot.Send(msg)
		if err != nil {
			println("出错")
			panic(err)
		}
	}
}

func Run(token string, webhookBase string, addr string, port string, debug bool) {

	// gin router
	router := gin.Default()

	// telegram
	initTelegram(token, webhookBase, debug)
	router.POST("/"+token, webhookHandler)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
