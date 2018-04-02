package utils

import (
	"encoding/json"
	"net/http"
)

func BuildJsonResponse(status string, writer http.ResponseWriter) {
	response, _ := json.Marshal(map[string]string{"status": status})

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(response)
}