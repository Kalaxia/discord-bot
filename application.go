package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"discord-bot/controller"
	"os"
	"net/http"
	"strings"
	"bytes"
	"log"
)

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
	app.Config.Bot.Channels = make(map[string]string, 0)
	app.Config.Bot.Channels["announcements"] = os.Getenv("DISCORD_ANNOUNCEMENTS_CHANNEL")

	// Initialize discord bot
	app.Session, err = discordgo.New("Bot " + app.Config.Bot.Token)
	if err != nil {
		log.Fatal("failed to create discord session")
	}
	err = app.Session.Open()
	if err != nil {
		log.Fatal("failed to open discord session")
	}
	log.Println("discord bot is running")
	app.Session.AddHandler(app.DiscordMessageHandler)

	// Initialize http router
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/polls/new", app.NewPollActionWrapper).Methods("POST")
}

// Run http server
func (app *Application) Run() {
	log.Println("http server is running")
	log.Fatal(http.ListenAndServe(app.Config.Http.Address, app.Router))
}

/*  
 * Define wrappers to avoid global discord session (global is ugly lulz)
 * **********************************************************************
 */

func (app *Application) NewPollActionWrapper(writer http.ResponseWriter, request *http.Request) {
	controller.NewPollAction(app.Session, app.Config.Bot.Channels["announcements"], writer, request)
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

		// Send response
		session.ChannelMessageSend(message.ChannelID, response.String())
	}
}
