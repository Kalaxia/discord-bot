package controller

import (
	"github.com/bwmarrin/discordgo"
	"encoding/json"
	"net/http"
	"discord-bot/utils"
	"bytes"
)

func NewPollAction(session *discordgo.Session, pollChannel string, writer http.ResponseWriter, request *http.Request) {
	var payload map[string]string

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", writer)
		return
	}

	// Build response
	pollUrl := "https://www.kalaxia.com/polls/" + payload["id"]

	var response bytes.Buffer
	response.WriteString("@everyone\n**Un nouveau vote à été soumis ! Allez voter sur ")
	response.WriteString(pollUrl)
	response.WriteString(" :D**")

	// Send discord message
	_, err := session.ChannelMessageSend(pollChannel, response.String())
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", writer)
}
