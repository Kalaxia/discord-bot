package main

import (
	"github.com/bwmarrin/discordgo"
	"os/signal"
	"syscall"
	"strings"
	"bytes"
	"log"
	"os"
)

var token = ""

func main() {
	// Create new session
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("failed to create session")
		return
	}
	log.Println("bot connected")

	// Add a message handler
	session.AddHandler(messageHandler)

	// Start session
	err = session.Open()
	if err != nil {
		log.Fatal("failed opening connection")
		return
	}
	log.Println("bot started")

	// Wait for interrupt (Ctrl +C)
	signalCanal := make(chan os.Signal, 1)
	signal.Notify(signalCanal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill) 
	<- signalCanal

	// Close the session
	session.Close()
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Check if message starts with "!tarvernier"
	if strings.HasPrefix(message.Content, "!tavernier") {
		// Split after one space
		drink := strings.SplitAfterN(message.Content, " ", 2)[1]
		if message.Author.ID == "321282244431970306" {
			drink = "de la souillure de chaussette"
		}
		
		// Build response
		var response bytes.Buffer
		response.WriteString("*donne ")
		response.WriteString(drink)
		response.WriteString(" à <@")
		response.WriteString(message.Author.ID)
		response.WriteString(">*")

		// Send response
		session.ChannelMessageSend(message.ChannelID, response.String())
	}

}
