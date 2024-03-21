package main

import (
	"log"
	"os"

	bot "figuriste.com/disgo-agt/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("DISCORD_BOT_TOKEN")

	bot.BotToken = botToken
	bot.Run() // call the run function of bot/bot.go
}
