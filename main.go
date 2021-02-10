package main

func main() {
	app := NewAppServer("")
	app.Middlewares()
	app.Routers()
	app.Listening()
}
