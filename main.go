package main

func main() {
	app := NewAppServer("8080")
	app.Middlewares()
	app.Routers()
	app.Listening()
}
