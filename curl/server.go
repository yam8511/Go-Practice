package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// req.ParseForm()
		if req.Method == "GET" || req.Method == "POST" {
			var v map[string]interface{}
			err := json.NewDecoder(req.Body).Decode(&v)
			if err != nil {
				// handle error
			}
			fmt.Println("header", req.Header)
			fmt.Println("json", v)
			fmt.Println(req.ContentLength)
			// firstname := req.FormValue("firstname")
			// lastname := req.FormValue("lastname")
			firstname := v["firstname"]
			lastname := v["lastname"]
			w.Write([]byte(fmt.Sprintf("[%s] Hello, %s %s!", req.Method, firstname, lastname)))
		} else {
			http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	}
}
