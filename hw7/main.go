package main

import (
	"fmt"
	"github.com/freakkid/Service-Computing/hw7/client"
	"time"
)

func syncTest() {
	start := time.Now()
	
	x := client.HTTPGet("kid")
	y := client.HTTPGet("sherry")
	z := client.HTTPGet("bingo")
	m := client.HTTPGet("lock")
	n := client.HTTPGet("exit")

	fmt.Printf("%s", x+y+z+m+n)
	fmt.Println("Time:", time.Since(start))
}

func asyncTest() {
	s := make(chan string)
	start := time.Now()

	go client.HTTPGetAsync("kid", s)
	go client.HTTPGetAsync("sherry", s)
	go client.HTTPGetAsync("bingo", s)
	go client.HTTPGetAsync("lock", s)
	go client.HTTPGetAsync("exit", s)

	fmt.Printf("%s", <-s+<-s+<-s+<-s+<-s)
	fmt.Println("Time:", time.Since(start))
}

func main() {
	fmt.Printf("%s\n\n", "使用 go HTTPClient 实现图 6-2 的 Naive Approach")
	syncTest()

	fmt.Println("------------------------------------")

	fmt.Printf("%s\n\n", "为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3")
	asyncTest()

	fmt.Println("\n[本次作业使用了百度翻译API@百度翻译]")
}
