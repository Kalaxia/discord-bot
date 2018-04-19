package main

import(
	"discord-bot/controller"
	"discord-bot/server"
	"github.com/gorilla/mux"
)

func main() {
	server.App.Initialize()
	InitializeRouter()
	server.App.Run()
}

func InitializeRouter() {
	server.App.Router = mux.NewRouter()
	server.App.Router.HandleFunc("/polls/new", controller.AddPollAction).Methods("POST")
	server.App.Router.HandleFunc("/tickets/new", controller.AddTicketAction).Methods("POST")
	server.App.Router.HandleFunc("/tickets/update", controller.UpdateTicketAction).Methods("POST")
	server.App.Router.HandleFunc("/tickets/delete", controller.RemoveTicketAction).Methods("POST")
}
