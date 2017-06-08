package main

import (
	// "database/sql"

	// "encoding/json"
	"fmt"
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

type UserData struct {
	Name  string
	email string
}

type UserDataArray struct {
	User []UserData
}

// Zuzu : App
var Zuzu App

// func (app App) dbHandle(handle func(*sql.DB)) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("Error: ", err)
// 		}
// 	}()
//
// 	// 建立連線
// 	db, err := connect("mysql", "go", "root", "a7319779")
// 	// 於結束前, 關閉連線
// 	defer db.Close()
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	handle(db)
// }

// func connect(driver string, database string, username string, password string) (*sql.DB, error) {
// 	// "root:a7319779@/go?charset=utf8"
// 	info := username + ":" + password + "@/" + database + "?charset=utf8"
// 	db, err := sql.Open(driver, info)
// 	return db, err
// }

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	Zuzu.startServe()

	// jsonString := `{"Age":23,"Lang":["Go","PHP",219],"Name":"Zuolar"}`
	// jsonData := []byte(jsonString)
	// var data map[string]interface{}
	// err := json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println("Json Decode Error", err)
	// } else {
	// 	fmt.Println(data)
	// }
	//
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	var v map[string]interface{}
	// 	err := json.NewDecoder(r.Body).Decode(&v)
	// 	if err != nil {
	// 		fmt.Println("Json Decode Error", err)
	// 	} else {
	// 		fmt.Println(data)
	// 	}
	// })
}

func (app App) startServe() {
	// 當瀏覽器輸入根目錄時會呼叫 indexHandle() 涵式
	// 當 Request 找不到指定 url 會選擇根目錄
	// 若連 根目錄 也找不到, 會跑出 "404 page not found"
	http.HandleFunc("/", serverHandler)
	http.ListenAndServe(":8000", nil)
}

func serverHandler(res http.ResponseWriter, req *http.Request) {
	// 一旦有例外發生, 顯示錯誤
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprint(res, err)
		}
	}()

	request := &Request{self: req, path: req.URL.EscapedPath()}
	// fmt.Println("input rr", rr.allInput())

	Zuzu.req = request
	Zuzu.res = res
	// Zuzu = &App{res: res, req: req}
	Zuzu.Start()
}

// Start : 處理Request
func (app App) Start() {
	currentURL := app.req.URL.EscapedPath()
	fmt.Println("path", currentURL)
	route := ServerRoute()
	handler, exists := route[currentURL]

	if !exists {
		app.view("404", nil)
	} else {
		handler(&app)
	}
}
