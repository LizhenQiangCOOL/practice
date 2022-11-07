package main

import (
	"fmt"
	"time"
)

func main() {
	// Ticker心脏、心跳，会一直间隔输出  <-chan
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool) //zero channel

	go func() {
		for {
			select {
			//往zero channle的goroutine等待队列加入该goroutine
			case result := <-done:
				fmt.Println("result:", result)
				return
				// 等待ticker的chan有值
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() //自己调用Stop
	done <- true
	fmt.Println("Ticker stopped")

	//为了查看 另外一个go程的结果
	time.Sleep(1600 * time.Millisecond)
}
