package controller

import (
	"net/http"
	"discord-bot/exception"
	"discord-bot/server"
	"discord-bot/utils"
	"strconv"
)

func AddPollAction(writer http.ResponseWriter, request *http.Request) {
	defer utils.CatchException(writer)
	payload := utils.ParseJsonRequest(request)
	if _, ok := payload["id"]; !ok {
		panic(exception.New(400, "Poll ID is missing", nil))
	}
	server.SendDiscordMessage(
		"board",
		"@everyone **Un nouveau vote à été soumis ! Allez voter sur https://www.kalaxia.com/polls/" +
		strconv.FormatInt(int64(payload["id"].(float64)), 10) + "** :envelope_with_arrow:",
	);
	utils.SendResponse(writer, 204, "")
}
