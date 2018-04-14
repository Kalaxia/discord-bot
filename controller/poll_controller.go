package controller

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"discord-bot/utils"
	"strconv"
	"log"
)

func PollAddAction(session *discordgo.Session, pollChannel string, writer http.ResponseWriter, request *http.Request) {
	/* payload, err := utils.ParseJsonRequest(request) */
	var payload map[string]int
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	if _, ok := payload["id"]; !ok {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	}
	// Build response
	pollUrl := "https://www.kalaxia.com/polls/" + string(payload["id"])

	var response string
	response = "@everyone\n**Un nouveau vote à été soumis ! Allez voter sur " + pollUrl + " :D**"

	// Send discord message
	_, err := session.ChannelMessageSend(pollChannel, response)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", "notification sent", writer)
}
