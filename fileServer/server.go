package main

import (
	"net/http"
)

func main() {
	// 第一個參數為route, 第二個參數為callback func
	// http.FileServer 建立一個 檔案伺服器
	// http.Dir("檔案資料夾路徑")
	http.Handle("/", http.FileServer(http.Dir("/home/zuolar/Go-Practice")))

	// 開啟伺服器並監聽Port
	http.ListenAndServe(":80", nil)
}
