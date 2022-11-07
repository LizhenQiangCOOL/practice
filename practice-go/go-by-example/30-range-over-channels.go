package main

import "fmt"

func main() {

	queue := make(chan string, 2) //buffer channel
	queue <- "one"
	queue <- "two"
	close(queue) //关闭了chan就只可以读了

	for elem := range queue {
		fmt.Println(elem)
	}

}
