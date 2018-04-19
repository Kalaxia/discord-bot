package controller

import (
	"net/http"
	"discord-bot/server"
	"discord-bot/utils"
	"log"
)

func AddTicketAction(writer http.ResponseWriter, request *http.Request) {
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	if !utils.CheckKeys(payload, "status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	}
	_, err := server.App.Session.ChannelMessageSend(
		server.App.Config.Bot.Channels["board"],
		`La carte **` + payload["title"] + `** à été publiée (statut: **` + payload["status"] + `**) !
		(https://www.kalaxia.com/feedbacks/` + payload["slug"] + `)`,
	)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}
	utils.BuildJsonResponse("ok", "notification sent", writer)
}

func UpdateTicketAction(writer http.ResponseWriter, request *http.Request) {
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}
	if !utils.CheckKeys(payload, "old_status", "new_status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	}
	_, err := server.App.Session.ChannelMessageSend(
		server.App.Config.Bot.Channels["board"],
		"La carte **" + payload["title"] + "** est passé de **" + payload["old_status"] + "** à **" + payload["new_status"] + `**
		 (https://www.kalaxia.com/feedbacks/` + payload["slug"] + ")",
	)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}
	utils.BuildJsonResponse("ok", "notification sent", writer)
}

func RemoveTicketAction(writer http.ResponseWriter, request *http.Request) {
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}
	if !utils.CheckKeys(payload, "status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	}
	_, err := server.App.Session.ChannelMessageSend(
		server.App.Config.Bot.Channels["board"],
		"La carte **" + payload["title"] + "** (statut: **" + payload["status"] + `**) à été supprimé !
		(https://www.kalaxia.com/feedbacks/` + payload["slug"] + ")",
	)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}
	utils.BuildJsonResponse("ok", "notification sent", writer)
}
