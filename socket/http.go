package main

import (
	"fmt"
	"github.com/yam8511/ZuZuGo"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello World")

	var app = new(ZuZuGo.App)
	app.SetRoute(map[string]func(*ZuZuGo.App){
		"/": indexHandler,
	})
	app.StartServe(8000)
}

func indexHandler(app *ZuZuGo.App) {
	msg := app.Req.Input("msg")
	data := ""
	if msg != nil {
		str := []byte(fmt.Sprint(msg))
		err := ioutil.WriteFile("aa.txt", str, 0777)
		if err != nil {
			fmt.Println("error:", err)
		}
	}

	content, err := ioutil.ReadFile("aa.txt")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		data = string(content)
	}

	app.Res.View("./index", data)
}
