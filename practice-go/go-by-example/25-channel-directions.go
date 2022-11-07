package main

import "fmt"

// chan<-  往chan里放东西
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan 往chan里拿东西
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
