package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string
var PrivateChannelID string

func Run() {
	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error initiating Discord connection:", err)
	}

	// add a event handler
	discord.AddHandler(receiveMessage)

	// open session
	discord.Open()
	defer discord.Close() // close session, after function termination

	// keep bot running untill there is NO os interruption (ctrl + C)
	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func receiveMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// prevent bot responding to its own message
	if (message.Author.ID == discord.State.User.ID) ||
		(message.ChannelID != PrivateChannelID && !isInList(message.Mentions, discord.State.User)) {
		return
	}

	// respond to user message
	switch {
	case strings.HasPrefix(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "This _!help_ command is not yet implemented")
	default:
		result, err := GetCompletion(message.Author.Username, message.ContentWithMentionsReplaced())
		if err != nil {
			fmt.Println(result, err)
			discord.ChannelMessageSendTTS(message.ChannelID, "Sorry there's a problem with my model.")
		} else {
			discord.ChannelMessageSend(message.ChannelID, result)
		}
	}
}

func isInList(users []*discordgo.User, targetUser *discordgo.User) bool {
	for _, user := range users {
		if user.Mention() == targetUser.Mention() {
			return true
		}
	}
	return false
}
