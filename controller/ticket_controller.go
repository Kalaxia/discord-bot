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
	if !utils.CheckKeys(payload, "status", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"La carte **" + payload["title"].(string) + "** (statut: **" + payload["status"].(string) + "**) à été supprimée ! :wastebasket:",
	)
	utils.SendResponse(writer, 204, "")
}

func AssignToTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "developper", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"**" + payload["developper"].(string) + "** à été assigné sur la carte **" + payload["title"].(string) + "**" +
		" (htts://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + ") :construction_site:",
	)
	utils.SendResponse(writer, 204, "")
}

func ValidateTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "tester", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"**" + payload["tester"].(string) + "** à validé la carte **" + payload["title"].(string) + "**" +
		" (htts://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + ") :white_check_mark:",
	)
	utils.SendResponse(writer, 204, "")	
}

func CommentTicketAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if !utils.CheckKeys(payload, "author", "slug", "title") {
		panic(exception.New(400, "Invalid data", nil))
	}
	server.SendDiscordMessage(
		"board",
		"**" + payload["author"].(string) + "** à commenté sur la carte **" + payload["title"].(string) + "**" +
		" (htts://www.kalaxia.com/feedbacks/" + payload["slug"].(string) + ") :writing_hand:",
	)
	utils.SendResponse(writer, 204, "")
}
