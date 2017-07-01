package main

import (
	// "fmt"
	"log"
	"net/http"

	socket "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socket.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", socketConnection)
	server.On("error", socketErrorHandler)

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
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
