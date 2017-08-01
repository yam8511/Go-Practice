package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	// 開啟伺服器並監聽Port
	fmt.Println("Start Serve on localhost:80")
	http.ListenAndServe(":80", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// 開啟瀏覽器, 輸入 localost:8000
	// 可以看到 Hello World
	fmt.Fprintln(res, "Hello World")
}
