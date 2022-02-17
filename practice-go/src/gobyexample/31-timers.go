package main

import (
	"fmt"
	"time"
)

func main()  {

	timer1 := time.NewTimer(5 * time.Second) //return  <-chan Time

	fmt.Println("Timer 1 not start")
	// 等待定时器结束，在chan的goroutine等待队列加入当前的goroutinue，当chan一旦有值，该goroutinue就会处理
	<- timer1.C
	fmt.Println("Timer 1 fired")

	timer2 :=  time.NewTimer(time.Second) //return <-chan current time
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 自己调用，停止timer2定时器，成功停止返回True
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stopped")
	}

	// 如果只是停止
	time.Sleep(2 * time.Second)
}
