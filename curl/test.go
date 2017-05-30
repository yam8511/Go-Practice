package main

import (
	// "bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var jsonStr string = `{"firstname": "Zuolar","lastname": "Lee"}`

	url := "http://localhost:8000"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
