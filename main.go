package main

import (
	"bot_steamDB/src/clog"
	"bot_steamDB/src/consts"
	"bot_steamDB/src/handler_error"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		herr.HandlerError(err, "No .env file found")
	}
	apiToken := os.Getenv(consts.NAME_API_TOKEN)
	bot, err := tgBotApi.NewBotAPI(apiToken)
	herr.HandlerFatal(err, "Token not correct")

	bot.Debug = true

	clog.Logger.Info("[BOT]", "Authorized on account", bot.Self.UserName)

	Listen(bot)
}

// Listen
// This code listens for incoming messages in a chat bot API by creating an updates channel and continuously checking
// for updates. If the update received is a message,
// the code creates a new message using the received message's chat ID and text and sends it using the bot's send
// function. Any errors that occur during the send process are logged.
func Listen(bot *tgBotApi.BotAPI) {
	updateConfig := tgBotApi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	// This code receives updates from a chat bot API and checks if the update is a message. If it is,
	// the code creates a new message using the update's chat ID and text,
	// then sends the message using the bot's send function.
	// The code logs any errors that occur during the send process.
	for update := range updates {
		if update.Message != nil { // If we got a message
			clog.Logger.Info("[BOT MESSAGE]", update.Message.From.UserName, update.Message.Text)

			msg := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := bot.Send(msg)
			herr.HandlerError(err, "[BOT] Not sent message")
		}
	}
}
