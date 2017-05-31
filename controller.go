package main

import (
	"crypto/md5"
	_ "database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	_ "net/url"
	"strconv"
	"text/template"
	"time"
)

func socketHandler(app *App) {
	data := make(map[string]interface{})
	data["Host"] = app.req.Host
	data["Port"] = "8000"
	app.view("socket", data)
}

func loginHandler(app *App) {
	res, req := app.res, app.req

	// 讀取 Cookie 值
	cookieValue, _ := req.Cookie("uid")
	fmt.Println("cookie", cookieValue)
	fmt.Println("all cookie", req.Cookies())

	// 設置 Cookie 值
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "uid", Value: "Zuolar", Expires: expiration}
	http.SetCookie(res, &cookie)

	// v := url.Values{}
	// v.Set("name", "Ava")
	// v.Add("friend", "Jess")
	// v.Add("friend", "Sarah")
	// v.Add("friend", "Zoe")
	// fmt.Fprintln(app.res, v.Encode())
	// fmt.Println(v.Get("name"))
	// fmt.Println(v.Get("friend"))
	// fmt.Println(v.Get("friend"))
	// fmt.Println(v["friend"])
}

func indexHandler(app *App) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	app.view("index", token)

	app.req.ParseForm()

	// fmt.Println("username:", template.HTMLEscapeString(app.req.Form.Get("item"))) //输出到服务器端
	// fmt.Println("password:", template.HTMLEscapeString(app.req.Form.Get("price")))
	// template.HTMLEscape(app.res, []byte(app.req.Form.Get("item")))

	// item1 := app.req.Form.Get("item")
	// price1 := app.req.Form["price"]
	// fmt.Println("form", app.req.Form, item1, price1)
	// fmt.Println("path", app.req.URL.Path)
	// fmt.Println("Scheme", app.req.URL.Scheme)

	item := template.HTMLEscapeString(app.req.FormValue("item"))
	price := template.HTMLEscapeString(app.req.FormValue("price"))
	var p Money = Money{Item: item, Price: price} // 要送到前端的訊息
	app.view("price", p)
}

func priceHandler(app *App) {
	app.req.ParseForm()
	item := app.req.FormValue("item")
	price := app.req.FormValue("price")
	data := make(map[string]interface{})
	data["Item"] = item
	data["Price"] = price
	app.view("price", data)
}

func sqlHandler(app *App) {
	app.req.ParseForm()
	item := app.req.FormValue("item")
	price := app.req.FormValue("price")
	data := make(map[string]interface{})
	data["Item"] = item
	data["Price"] = price
	app.view("price", data)
}
