package main

import (
	"fmt"
	"time"
)

func main(){

	c1 := make(chan string, 1)

	go func(){
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <- c1:
		fmt.Println(res)

		//使用 time.After 来做超时限制
		//func After(d Duration) <-chan Time
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	//再超时限制内，chan产生了值
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")

	}
}