package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func runForever(id int) {
	fmt.Printf("我是佳菁 %d 號\n", id)
	for {
		runtime.Gosched()
	}
}

func wg() {
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go runForever(i)
	}
	wg.Wait()
}
func main() {
	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	wg()
}
