package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	runtime "runtime"
	_ "strings"
	time "time"
	// zz "zuolar"
)

type Per struct {
	name  string
	data  interface{}
	intro func()
}

type Auth struct {
	*Per
	email string
}

type Handler interface {
	process()
}

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s.Servers[0].ServerName)
}

func main1() {
	file, err := os.Open("index.html")
	if err != nil || os.IsExist(err) {
		fmt.Println("error", err)
	} else {
		fmt.Println("file", file)
	}

	// mm := make(map[string]int)
	// mm["a"] = 12
	// mm["b"] = 34
	// fmt.Println(mm, mm["a"])

	// type person struct {
	// 	name string
	// 	age  int
	// }
	//
	// var snoopy person = person{name: "Snoopy", age: 100}
	// me := person{name: "Zuolar", age: 23}
	// fmt.Println(snoopy, me, me.name) // ---> {Snoopy 100} {Zuolar 23} Zuolar
	//
	// var a interface{}
	// a = 123
	// fmt.Println(a) // ---> 123
	// a = "hello"
	// fmt.Println(a) // ---> hello
	// a = true
	// fmt.Println(a) // ---> true

	// ci := make(chan int)
	// ci <- 1
	// fmt.Println(<-ci)
	// ci <- 2
	// fmt.Println(<-ci)
	// ci <- 3
	// fmt.Println(<-ci)
	// cs := make(chan string)
	// cf := make(chan interface{})
	// a := []int{1, 2, 3, 4, 5}
	// go sum(a[:len(a)/2], ci)
	// go sum(a[len(a)/2:], ci)
	// x, y := <-ci, <-ci
	// fmt.Println(x, y, x+y, ci)
	// c := make(chan int, 10)
	// go fibonacci(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	// ca := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-ca)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci2(ca, quit)
}

func timeout() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

func fibonacci2(ca, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case ca <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func rungo() {
	runtime.GOMAXPROCS(2)
	go say("go")
	go say("hello")
	go say("Zuolar")
	go say("Google")
	say("Jia Jing")
}

func say(s string) {
	count := 5
	for i := 0; i < count; i++ {
		runtime.Gosched()
		// runtime.NumCPU()
		fmt.Println(s)
	}
}

func defConstAndSwitch() {
	const (
		a, x = iota, iota
		y    = iota
		z    = iota
	)

	const (
		b, c = iota, iota
		d    = iota
		f    = iota
	)

	fmt.Println(a, x, y, z)
	fmt.Println(b, c, d, f)
	// startServer()
	foo := Per{name: "Zuolar"}
	Handler.process(foo)

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func (p Per) process() {
	fmt.Println("Per | " + p.name)
}

func (p Auth) process() {
	fmt.Println("Auth | " + p.email)
}

func embed() {
	a := &Per{name: "Zuolar", data: "Google"}
	b := &Auth{Per: a, email: "yam8511@gmail.com"}
	b.Per.say()
}

func errorProcess() {
	if num, err := therror(1); err != nil {
		fmt.Println(num, err)
	} else {
		fmt.Println(num, err)
	}

	// 製作的套件不能有重複的 main func()
	// zz.Say()

	try(func() {
		foo(9)
	}, func(e interface{}) {
		fmt.Println(e)
	})

	// number := 11
	// try(func() {
	// 	if number < 10 {
	// 		panic("number is less than 10")
	// 	} else if number > 10 {
	// 		panic("number is greater than 10")
	// 	} else {
	// 		fmt.Println("ok")
	// 	}
	// }, func(e interface{}) {
	// 	fmt.Println(e)
	// })

}

func newClass() {
	b := &Per{name: "Zuolar"}
	b.say()
	b.intro = func() {
		fmt.Println("I am " + b.name)
	}

	b.intro()

	c := &Per{name: "AA", data: 123}
	c.say()

	a := newPer("CCC", []int{1, 2, 3})
	a.say()

	d := newPer("DDD", "Taiwan google")
	d.say()
}

func newPer(name string, data interface{}) (person Per) {
	person = Per{name: name, data: data}
	return person
}

func (p *Per) say() {
	fmt.Println("Hello, I am "+p.name, p.data)
}

func foo(number int) {
	if number < 10 {
		panic("number is less than 10")
	} else if number > 10 {
		panic("number is greater than 10")
	} else {
		fmt.Println("ok")
	}
}

func try(fn func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()

	fn()
}

func therror(number int) (int, error) {
	if number != 1 {
		return -1, errors.New("$number is not 1")
	}

	return number, nil
}

func hashArray() {
	data := map[string]interface{}{}
	data["username"] = "YamiOdymel"
	data["password"] = 123
	data["aa"] = "asd"
	username, exist := data["password"]
	if exist {
		fmt.Printf("%T: '%s'\n", username, username)
		fmt.Println(exist, username)
	} else {
		fmt.Printf("Not %T: '%s'\n", username, username)
		fmt.Println(exist)
	}
}

func demo() {
	// var a, b string = "aa", ""
	// c := "cc"
	// c = "dd"
	// b = "bb"
	// fmt.Println("ok", a, b, c)
	// d, e := fmt.Printf("a type: %T value: %s,\n b type: %T, value: %s\n c type: %T value: %s\n", a, a, b, b, c, c)
	//
	// fmt.Println(d, e)
	//
	// if e == nil {
	// nun, _ := fmt.Println("no")
	// fmt.Println("num", nun)
	// }
	//
	// // C := []int{1, 2, 3}
	// var C [3]int
	// C[0] = 456
	// fmt.Println(C[0:0])
	// _, time := test()
	// fmt.Println(time)
	//
	// aa := func() string {
	// return "hello"
	// }
	//
	// fmt.Println(aa())
	// var pp *int = &C[0]
	// C[0] = 123
	// fmt.Println(*pp)
	// fmt.Println(&C[0])
	//
	// p := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(p[0:5]) // 輸出：[1]
	// fmt.Println(p[1:2]) // 輸出：[]  （！注意這跟 PHP 不一樣！）
	// fmt.Println(p[1:])  // 輸出：[2, 3, 4, 5, 6]
	// fmt.Println(p[:3])  // 輸出：[1]
	//
	// data := make(map[string]int)
	// data["zz"] = 1
	// data["kk"] = 2
	// fmt.Println(data, data["zz"])
	// mixedData := make(map[string]interface{})
	// mixedData["username"] = "Zuolar"
	// mixedData["time"] = 219
	// goto aaa
	// defer fmt.Println(mixedData)
	//
	// aaa:
	// var mix interface{} = "aa"
	// mix = 123
	// fmt.Println(mix)
	// count := 10
	// for i := 0; i < count; i++ {
	// fmt.Println(i)
	// }

	// var data []string = []string{"a", "b", "c"}
	// for _, value := range data {
	// fmt.Println(value)
	// }

	// i := 0
	// for i < 10 {
	// i++
	// fmt.Println(i)
	// }
	// i := 0
	// for {
	// fmt.Println("WOW") // 輸出：WOWOWOWOWOWOWOWOW...
	// i++
	// if i > 3 {
	// break
	// }
	// }

	// data := "a,b,c,d,e"
	// arr := strings.Split(data, ",")
	// fmt.Println(arr[0])
	// fmt.Println(time.Now().Format("2006/1/2 3:4:5 Mon"))
}

func test() (string, int) {
	return "This is a test.", 123456
}
