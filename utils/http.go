package utils

import (
	"encoding/json"
	"net/http"
)

func BuildJsonResponse(status string, reason string, writer http.ResponseWriter) {
	response, _ := json.Marshal(map[string]string{"status": status, "reason": reason})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(response)
}