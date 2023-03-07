package core

import (
	"bot_steamDB/src/clog"
	"bot_steamDB/src/consts"
	herr "bot_steamDB/src/handler_error"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

var (
	Bot *tgBotApi.BotAPI
)

// ConfigBot
// This code configures a Telegram bot by retrieving the API token from an environment variable,
// creating a new Telegram Bot API with that token, setting the bot's debug mode to true,
// and logging a message indicating that the bot has been authorized. If the API token is not correct,
// it logs an error message.
func ConfigBot() {
	apiToken := os.Getenv(consts.NAME_API_TOKEN)
	bot, err := tgBotApi.NewBotAPI(apiToken)
	herr.HandlerFatal(err, "Token not correct")

	bot.Debug = true

	clog.Logger.Info("[BOT]", "Authorized on account", bot.Self.UserName)
}

// Listen
// This code listens for incoming messages in a chat bot API by creating an updates channel and continuously checking
// for updates. If the update received is a message,
// the code creates a new message using the received message's chat ID and text and sends it using the bot's send
// function. Any errors that occur during the send process are logged.
func Listen() {
	updateConfig := tgBotApi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := Bot.GetUpdatesChan(updateConfig)

	// This code receives updates from a chat bot API and checks if the update is a message. If it is,
	// the code creates a new message using the update's chat ID and text,
	// then sends the message using the bot's send function.
	// The code logs any errors that occur during the send process.
	for update := range updates {
		if update.Message != nil { // If we got a message
			clog.Logger.Info("[BOT MESSAGE]", update.Message.From.UserName, update.Message.Text)

			msg := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			_, err := Bot.Send(msg)
			herr.HandlerError(err, "[BOT] Not sent message")
		}
	}
}
