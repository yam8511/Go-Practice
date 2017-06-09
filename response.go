package main

import (
	"encoding/json"
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

	// json.Marshal 回傳 []byte 跟 error
	jsonData, err := json.Marshal(data)
	// 如果 err 不為 nil , 表示編碼有問題
	if err != nil {
		fmt.Fprint(res.self, "Json Encode Error")
	} else {
		// Json資料為 []byte 型態, 須轉字串才看得懂
		output := string(jsonData)
		fmt.Fprint(res.self, output)
		// ---> {"Age":23,"Lang":["Go","PHP",219],"Name":"Zuolar"}
		// ---> {"User":[{"Name":"Zuolar"},{"Name":"Golang"}]}
	}
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
