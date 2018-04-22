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
	server.App.Router.HandleFunc("/polls", controller.AddPollAction).Methods("POST")
	server.App.Router.HandleFunc("/tickets", controller.AddTicketAction).Methods("POST")
	server.App.Router.HandleFunc("/tickets", controller.UpdateTicketAction).Methods("PUT")
	server.App.Router.HandleFunc("/tickets", controller.RemoveTicketAction).Methods("DELETE")
	server.App.Router.HandleFunc("/tickets/comment", controller.CommentTicketAction).Methods("POST")
}
