package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request : 請求的物件
type Request struct {
	req  *http.Request
	path string
}

// input : Request取輸入參數的方法
func (this *Request) input() map[string]interface{} {
	fmt.Println("header", this.req.Header["Content-Type"])
	fmt.Println("body", this.req.Body)
	v := make(map[string]interface{})
	jsonData := make(map[string]interface{})
	// header, exists := this.req.Header["Content-Type"]

	err := this.req.ParseForm()
	if err == nil {
		for key, value := range this.req.Form {
			if len(value) == 1 {
				v[key] = value[0]
			} else {
				v[key] = value
			}
			fmt.Println("key: ", key, "value: ", value)
		}
	}

	err = json.NewDecoder(this.req.Body).Decode(&jsonData)
	if err == nil {
		for key, value := range jsonData {
			v[key] = value
		}
	}

	return v
}
