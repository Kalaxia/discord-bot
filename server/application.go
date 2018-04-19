package server

import (
	"discord-bot/exception"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"os"
	"net/http"
	"strings"
	"bytes"
	"log"
)

var App Application

type Application struct {
	Router		*mux.Router
	Session		*discordgo.Session
	Config 		ApplicationConfig
}

type HttpConfig struct {
	Address		string 				`yaml:"address"`
}

type BotConfig struct {
	Token		string
	Channels 	map[string]string
}

type ApplicationConfig struct {
	Bot 		BotConfig
	Http 		HttpConfig
}

// Configure bot & router
func (app *Application) Initialize() {
	var err error

	app.Config.Bot.Token = os.Getenv("DISCORD_SERVER_TOKEN")
	app.Config.Bot.Channels = map[string]string{
		"announcements": os.Getenv("DISCORD_ANNOUNCEMENTS_CHANNEL"),
		"board": os.Getenv("DISCORD_BOARD_CHANNEL"),
	}
	app.Session, err = discordgo.New("Bot " + app.Config.Bot.Token)
	if err != nil {
		panic(exception.New(500, "Discord session could not be created", err))
	}
	if err = app.Session.Open(); err != nil {
		panic(exception.New(500, "Discord session could not be opened", err))
	}
	log.Println("discord bot is running")
	app.Session.AddHandler(app.DiscordMessageHandler)
}

// Run http server
func (app *Application) Run() {
	log.Println("http server is running")
	log.Fatal(http.ListenAndServe(app.Config.Http.Address, app.Router))
}
/*
 * Define discord bot handlers
 */
func (app *Application) DiscordMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Check if message starts with "!tarvernier"
	if strings.HasPrefix(message.Content, "!tavernier") {
		var response bytes.Buffer

		// Split after one space
		parsedMessage := strings.SplitAfterN(message.Content, " ", 2)
		if len(parsedMessage) <= 1 {
			response.WriteString("*donne une chope vide à <@")
			response.WriteString(message.Author.ID)
			response.WriteString(">*")
		} else {
			drink := parsedMessage[1]

			if message.Author.ID == "321282244431970306" {
				drink = "de la souillure de chaussette"
			}

			response.WriteString("*donne ")
			response.WriteString(drink)
			response.WriteString(" à <@")
			response.WriteString(message.Author.ID)
			response.WriteString(">*")
		}
		session.ChannelMessageSend(message.ChannelID, response.String())
	}
}

func SendDiscordMessage(channel, message string) {
	if _, err := App.Session.ChannelMessageSend(
		App.Config.Bot.Channels[channel],
		message,
	); err != nil {
		panic(exception.New(500, "Message could not be sent to Discord", err))
	}
}
