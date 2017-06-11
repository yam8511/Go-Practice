package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	ZuZuGo "github.com/yam8511/ZuZuGo"
)

var app *ZuZuGo.App

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Hello World")
	app = new(ZuZuGo.App)
	route := map[string]func(*ZuZuGo.App){}
	route["/"] = indexHandler
	app.SetRoute(route)
	p, _ := strconv.Atoi(port)
	app.StartServe(uint(p))
}

func indexHandler(app *ZuZuGo.App) {
	app.Res.JSON("hello")
}
