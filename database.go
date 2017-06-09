package main

import (
	"database/sql"
	"fmt"
)

func (app App) dbHandle(handle func(*sql.DB)) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()

	// 建立連線
	db, err := connect("mysql", "go", "root", "a7319779")
	// 於結束前, 關閉連線
	defer db.Close()

	if err != nil {
		panic(err)
	}

	handle(db)
}

func connect(driver string, database string, username string, password string) (*sql.DB, error) {
	// "root:a7319779@/go?charset=utf8"
	info := username + ":" + password + "@/" + database + "?charset=utf8"
	db, err := sql.Open(driver, info)
	return db, err
}
