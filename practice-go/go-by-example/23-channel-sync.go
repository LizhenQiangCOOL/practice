package main

import (
	"fmt"
	"time"
)

// chan 同步
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//给chan输入一个值
	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	//等待 chan中的值
	<-done
	fmt.Println("主go程完成任务")
}
