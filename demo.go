package main

import (
	"fmt"

	ZuZuGo "github.com/yam8511/ZuZuGo"
)

var app *ZuZuGo.App

func main() {
	fmt.Println("Hello World")
	app = new(ZuZuGo.App)
	route := map[string]func(*ZuZuGo.App){}
	route["/"] = indexHandler
	app.SetRoute(route)
	app.StartServe(8000)
}

func indexHandler(app *ZuZuGo.App) {
	app.Res.JSON("hello")
}
