package utils

import (
	"discord-bot/exception"
	"encoding/json"
	"log"
	"net/http"
)

func SendJsonResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		panic(exception.New(500, "JSON response could not be encoded", err))
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(response)
}

func SendResponse(writer http.ResponseWriter, statusCode int, response string) {
	writer.WriteHeader(statusCode)
	writer.Write([]byte(response))
}

func ParseJsonRequest(request *http.Request) map[string]interface{} {
	payload := make(map[string]interface{})
    if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		panic(exception.New(400, "Request body could not be parsed into JSON", err))
	}
	return payload
}

func CatchException(w http.ResponseWriter) {
    r := recover()
    if r == nil {
        return
    }
    if exception, ok := r.(*exception.Exception); ok {
        message := ""
        if exception.Error != nil {
            message = "; [Error]: " + exception.Error.Error()
        }
        if exception.Message != "" || message != "" {
            log.Println("[Exception]: " + exception.Message + message)
        }
		SendJsonResponse(w, exception.Code, exception)
        return
    }
    if err, ok := r.(error); ok {
        log.Println("[Error]: " + err.Error())
        return
    }
    panic(r)
}
