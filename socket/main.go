package main

import (
	"fmt"
	"log"
	"net/http"

	socket "github.com/googollee/go-socket.io"
	"github.com/yam8511/ZuZuGo"
)

func main() {
	fmt.Println("Hello World")
	server, err := socket.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	} else {
		server.On("connection", socketConnection)
		server.On("error", socketErrorHandler)

		http.Handle("/socket.io/", server)
	}

	var app = new(ZuZuGo.App)
	app.SetRoute(map[string]func(*ZuZuGo.App){
		"/": indexHandler,
	})
	app.StartServe(8000)
}

// func indexHandler(app *ZuZuGo.App) {
// 	msg := app.Req.Input("msg")
// 	data := ""
// 	if msg != nil {
// 		str := []byte(fmt.Sprint(msg))
// 		err := ioutil.WriteFile("aa.txt", str, 0777)
// 		if err != nil {
// 			fmt.Println("error:", err)
// 		}
// 	}
//
// 	content, err := ioutil.ReadFile("aa.txt")
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		data = string(content)
// 	}
//
// 	app.Res.View("./index", data)
// }

func indexHandler(app *ZuZuGo.App) {
	if app.Req.Self.Method == "POST" {
		if app.Req.Input("uid") == "" {
			app.Res.View("index", "請輸入暱稱")
		} else {
			cookieValue := app.Req.Input("uid")
			cookie := http.Cookie{Name: "uid", Value: cookieValue.(string)}
			http.SetCookie(app.Res.Self, &cookie)
			app.Res.View("demo", cookieValue)
		}
	} else {
		cookieValue, err := app.Req.Self.Cookie("uid")
		log.Println("cookie", cookieValue, err)
		if err == nil {
			app.Res.View("demo", cookieValue)
		} else {
			app.Res.View("index", nil)
		}
	}
}

func socketErrorHandler(so socket.Socket, err error) {
	log.Fatal("error:", err)
}

func socketConnection(so socket.Socket) {
	log.Println("on connection")
	so.Join("chat")
	so.On("chat message", func(msg interface{}) {
		log.Println(msg)
		log.Println("emit:", so.Emit("chat message", msg))
		so.BroadcastTo("chat", "chat message", msg)
	})
	// Socket.io acknowledgement example
	// The return type may vary depending on whether you will return
	// For this example it is "string" type
	so.On("chat message with ack", func(msg string) string {
		return msg + " by Zuolar"
	})
	so.On("disconnection", func(user string) {
		log.Println("on disconnect", user)
	})
}
