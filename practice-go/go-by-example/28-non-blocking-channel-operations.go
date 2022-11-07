package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
		//和case关联的操作是阻塞的，default就会被执行
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	//zero channle 就是阻塞型channle，没有缓存的地方
	//往空channel发送值，会阻塞
	// 如果channel是缓存channel，则往里发送值，不会阻塞
	//尝试 message:=make(chan string) => message:=make(chan string, 1)
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
