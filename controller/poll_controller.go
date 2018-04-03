package controller

import (
	"github.com/bwmarrin/discordgo"
	"encoding/json"
	"net/http"
	"discord-bot/utils"
	"strconv"
	"bytes"
	"log"
)

func NewPollAction(session *discordgo.Session, pollChannel string, writer http.ResponseWriter, request *http.Request) {
	var payload map[string]interface{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	if !utils.CheckKey(payload, "id") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	}

	// Convert parameters
	pollId := strconv.FormatFloat(payload["id"].(float64), 'f', 0, 64)

	// Build response
	pollUrl := "https://www.kalaxia.com/polls/" + pollId

	var response bytes.Buffer
	response.WriteString("@everyone\n**Un nouveau vote à été soumis ! Allez voter sur ")
	response.WriteString(pollUrl)
	response.WriteString(" :D**")

	// Send discord message
	_, err := session.ChannelMessageSend(pollChannel, response.String())
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", "notification sent", writer)
}
