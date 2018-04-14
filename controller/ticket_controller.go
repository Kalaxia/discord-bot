package controller

import (
	"github.com/bwmarrin/discordgo"
	"net/http"
	"discord-bot/utils"
	"log"
)

func TicketAddAction(session *discordgo.Session, boardChannel string, writer http.ResponseWriter, request *http.Request) {
	// Parse json
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	// Check parameters
	if !utils.CheckKeys(payload, "status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	} 

	// Get parameters
	ticketStatus := payload["status"]
	ticketSlug := payload["slug"]
	ticketTitle := payload["title"]

	ticketUrl := "https://www.kalaxia.com/feedbacks/" + ticketSlug

	// Build response
	var response string
	response = "La carte **" + ticketTitle + "** à été publiée (statut: **" + ticketStatus + "**) ! "
	response = response + " (" + ticketUrl + ")"
	
	// Send discord message
	_, err := session.ChannelMessageSend(boardChannel, response)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", "notification sent", writer)
}

func TicketUpdateAction(session *discordgo.Session, boardChannel string, writer http.ResponseWriter, request *http.Request) {
	// Parse json
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	// Check parameters
	if !utils.CheckKeys(payload, "old_status", "new_status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	} 

	// Get parameters
	oldStatus := payload["old_status"]
	newStatus := payload["new_status"]
	ticketSlug := payload["slug"]
	ticketTitle := payload["title"]

	ticketUrl := "https://www.kalaxia.com/feedbacks/" + ticketSlug

	// Build response
	var response string
	response = "La carte **" + ticketTitle + "** est passé de **" + oldStatus + "** à **" + newStatus + "**"
	response = response + " (" + ticketUrl + ")"
	
	// Send discord message
	_, err := session.ChannelMessageSend(boardChannel, response)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", "notification sent", writer)
}

func TicketRemoveAction(session *discordgo.Session, boardChannel string, writer http.ResponseWriter, request *http.Request) {
	// Parse json
	var payload map[string]string
	if err := utils.ParseJsonRequest(request, &payload); err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "payload decoding error", writer)
		return
	}

	// Check parameters
	if !utils.CheckKeys(payload, "status", "slug", "title") {
		utils.BuildJsonResponse("error", "invalid payload", writer)
		return
	} 

	// Get parameters
	ticketStatus := payload["status"]
	ticketSlug := payload["slug"]
	ticketTitle := payload["title"]

	ticketUrl := "https://www.kalaxia.com/feedbacks/" + ticketSlug

	// Build response
	var response string
	response = "La carte **" + ticketTitle + "** (statut: **" + ticketStatus + "**) à été supprimé ! "
	response = response + " (" + ticketUrl + ")"
	
	// Send discord message
	_, err := session.ChannelMessageSend(boardChannel, response)
	if err != nil {
		log.Println(err)
		utils.BuildJsonResponse("error", "discord send message error", writer)
		return
	}

	// Send http response
	utils.BuildJsonResponse("ok", "notification sent", writer)
}