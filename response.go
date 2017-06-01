package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Response : 回應的物件
type Response struct {
	self http.ResponseWriter
}

// json : Response 回應 Json 資料
func (res *Response) json(data interface{}) {
	res.self.Header().Add("Content-Type", "application/json")
}

func (res *Response) view(page string, data interface{}) {
	prefix := "view/"
	targetPage := prefix + page + ".gtpl"
	t, err := template.ParseFiles(targetPage)
	if err != nil {
		fmt.Println("view error", err)
		panic("page didn't exist: " + page + ".gtpl")
	} else {
		t.Execute(res.self, data)
	}
}
