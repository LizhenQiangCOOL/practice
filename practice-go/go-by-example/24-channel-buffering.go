package main

import "fmt"

func main() {

	// 若为不填数值2，则为zero channel，会阻塞，等待接受的go程
	// zero channel 会引起 data rate
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
