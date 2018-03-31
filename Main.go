package main

func main() {
	var app Application
	app.Initialize("config.yml")
	app.Run()
}