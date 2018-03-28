package main

// var token = "NDA4MDE2MTE1MTQ3NDcyODk2.DZ12iA.iyVz3YqSkFlNQ0kswXB3S7cOg98"

func main() {
	var app Application
	app.Initialize("config.yml")
	app.Run()
}