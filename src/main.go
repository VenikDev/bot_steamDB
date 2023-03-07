package main

import (
	"bot_steamDB/src/core"
	"bot_steamDB/src/handler_error"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		herr.HandlerFatal(err, "No .env file found")
	}

	core.ConfigBot()
	core.Listen()
}
