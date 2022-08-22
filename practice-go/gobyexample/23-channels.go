package main

import "fmt"

func main() {

	// chan 的定义，make(chan  类型)
	// zero channel
	messages := make(chan string)

	go func() { messages <- "ping"}()


	msg := <-messages
	fmt.Println("I'm waiting")

	fmt.Println(msg)
}