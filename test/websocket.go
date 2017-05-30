package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"text/template"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, reply); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/chat", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println("view error", err)
			panic("page didn't exist: index.html")
		} else {
			t.Execute(res, nil)
		}
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
