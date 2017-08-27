package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range t1.C {
			fmt.Println("Tick at", t.String())
		}
	}()

	time.Sleep(time.Second * 5)
	t1.Stop()
	fmt.Println("Exit")
}
