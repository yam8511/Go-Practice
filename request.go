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

func (this *Request) input() map[string]interface{} {
	fmt.Println("header", this.req.Header["Content-Type"])
	fmt.Println("body", this.req.Body)
	v := make(map[string]interface{})
	header, exists := this.req.Header["Content-Type"]
	if exists && header[0] == "application/json" {
		err := json.NewDecoder(this.req.Body).Decode(&v)
		if err != nil {
			panic(err)
		}
	} else {
		err := this.req.ParseForm()
		if err != nil {
			panic(err)
		}

		for key, value := range this.req.Form {
			fmt.Println("key: ", key, "value: ", value)
		}
	}
	return v
}
