package main

import (
	"fmt"
	"time"
)

type Request struct {
	time int
}

var MaxOutstanding = 2
var sem = make(chan int, MaxOutstanding)

func handler(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func process(req *Request) {
	fmt.Printf("sleep %d(address: %p)s \r\n", *req, req)
	time.Sleep(time.Duration(req.time) * time.Second)
}

func ServerYes(clientRequests chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handler(clientRequests)
	}
	<-quit
}

func main() {
	handles := make(chan *Request, 5)
	quit := make(chan bool)
	// 任务写入，当没有任务时关闭 handles
	for i := 1; i <= 5; i++ {
		handles <- &Request{5}
	}
	close(handles)
	go ServerYes(handles, quit)
	select {
	// 获取到某个停止信号
	case <-time.After(10 * time.Second):
		quit <- true
	}
}
