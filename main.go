package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"strings"

	"github.com/RafaelYon/kkkj-bot/command"
	"github.com/RafaelYon/kkkj-bot/repository"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Carregando env")
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	devRepository *repository.Dev
)

func main() {
	LoadEnv()

	devRepository = &repository.Dev{}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		os.Exit(1)
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	if err = dg.Open(); err != nil {
		fmt.Println("error opening connection,", err)
		os.Exit(1)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if strings.HasPrefix(m.Content, ";points") {
		command.ContabilizePoints(m.Content, devRepository)
		command.PrintPoints(s, m, devRepository)
	}

	if strings.HasPrefix(m.Content, ";poll") {
		command.PrintPoints(s, m, devRepository)
	}

	if strings.HasPrefix(m.Content, ";oh dad") {
		s.ChannelMessageSend(m.ChannelID, "STOP 👉👈")
	}

	if strings.HasPrefix(m.Content, ";help") {
		command.Help(s, m)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == ";pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
