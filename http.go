package main

import (
	// "database/sql"
	"fmt"
	"strings"
	// _ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Money struct {
	Item  string
	Price string
}

// App : Http App
type App struct {
	req *Request
	res *Response
}

// Zuzu : App
var Zuzu App

func main() {
	Zuzu.startServe(8000)
}

func (app App) startServe(port uint) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// 當瀏覽器輸入根目錄時會呼叫 indexHandle() 涵式
	// 當 Request 找不到指定 url 會選擇根目錄
	// 若連 根目錄 也找不到, 會跑出 "404 page not found"
	http.HandleFunc("/", serverHandler)
	listenPort := fmt.Sprintf(":%d", port)
	http.ListenAndServe(listenPort, nil)
}

func serverHandler(res http.ResponseWriter, req *http.Request) {
	// 一旦有例外發生, 顯示錯誤
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprint(res, err)
		}
	}()

	// request := &Request{self: req, path: req.URL.EscapedPath()}
	// fmt.Println("input rr", rr.allInput())

	Zuzu.req = &Request{self: req, path: req.URL.EscapedPath()}
	Zuzu.res = &Response{self: res}
	// Zuzu = &App{res: res, req: req}
	Zuzu.Start()
}

// Start : 處理Request
func (app App) Start() {
	currentURL := app.req.self.URL.EscapedPath()
	fmt.Println("path", currentURL)
	route := ServerRoute()
	handler, exists := route[currentURL]

	if !exists {
		fmt.Println("Header", app.req.self.Header)
		headerAccept, headerExists := app.req.self.Header["Accept"]
		substr := "image"
		needResponse := true
		if headerExists {
			for _, acceptItem := range headerAccept {
				if strings.Contains(acceptItem, substr) {
					needResponse = false
					break
				}
			}
		}

		if needResponse {
			app.view("404", nil)
		}
	} else {
		handler(&app)
	}
}
