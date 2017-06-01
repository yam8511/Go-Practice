package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Request : 請求的物件
type Request struct {
	self *http.Request
	path string
}

// allInput : Request取所有輸入參數的方法
func (req *Request) allInput() map[string]interface{} {
	v := make(map[string]interface{})
	jsonData := make(map[string]interface{})

	err := req.self.ParseForm()
	if err == nil {
		for key, value := range req.self.Form {
			if len(value) == 1 {
				v[key] = value[0]
			} else {
				v[key] = value
			}
			fmt.Println("key: ", key, "value: ", value)
		}
	}

	err = json.NewDecoder(req.self.Body).Decode(&jsonData)
	if err == nil {
		for key, value := range jsonData {
			v[key] = value
		}
	}

	return v
}

// input : Request 取指定 key 值
func (req *Request) input(key string) interface{} {
	allInput := req.allInput()
	val, exists := allInput[key]
	if !exists {
		return nil
	}
	return val
}
