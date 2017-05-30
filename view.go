package main

import (
	"fmt"
	"text/template"
)

func (app *App) view(page string, data interface{}) {
	prefix := "view/"
	targetPage := prefix + page + ".gtpl"
	t, err := template.ParseFiles(targetPage)
	if err != nil {
		fmt.Println("view error", err)
		panic("page didn't exist: " + page + ".gtpl")
	} else {
		t.Execute(app.res, data)
	}
}

func (app *App) json(interface{}) *App {
	var json string
	fmt.Fprint(app.res, json)
	return app
}
