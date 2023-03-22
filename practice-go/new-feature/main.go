package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
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

func print111() {
	for i := 0; i < 1000; i++ {
		fmt.Println("111111" + "_ i is " + strconv.Itoa(i))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
func printgo() {
	go print111()
	for j := 0; j < 3; j++ {
		fmt.Println("jjjjj")
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen")
	c.L.Unlock()
}
