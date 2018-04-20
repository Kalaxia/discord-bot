package controller

import (
	"net/http"
	"discord-bot/exception"
	"discord-bot/server"
	"discord-bot/utils"
)

func AddTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "status", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"La carte **" + payload["title"].(string) + "** à été publiée (statut: **" + payload["status"].(string) + "**) ! " +
		"https://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + " :rocket:",
	)
	utils.SendResponse(writer, 204, "")
}

func UpdateTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "old_status", "new_status", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"La carte **" + payload["title"].(string) + "** est passée de **" + payload["old_status"].(string) + "** à **" + payload["new_status"].(string) + "**" +
		" https://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + " :rocket:",
	)
	utils.SendResponse(writer, 204, "")
}

func RemoveTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "status", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"La carte **" + payload["title"].(string) + "** (statut: **" + payload["status"].(string) + "**) à été supprimée !" +
		"https://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + " :wastebasket:",
	)
	utils.SendResponse(writer, 204, "")
}
