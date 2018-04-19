package controller

import (
	"net/http"
	"discord-bot/server"
	"discord-bot/utils"
	"log"
)

func AddPollAction(writer http.ResponseWriter, request *http.Request) {
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
	if _, err := server.App.Session.ChannelMessageSend(
		server.App.Config.Bot.Channels["board"],
		`@everyone\n**Un nouveau vote à été soumis !
		Allez voter sur https://www.kalaxia.com/polls/` + string(payload["id"]) + ` :D**`,
	); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}
	utils.BuildJsonResponse("ok", "notification sent", writer)
}
