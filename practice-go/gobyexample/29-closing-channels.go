package main

import "fmt"

func main()  {
	jobs := make(chan int, 5)  //buffer channel
	done := make(chan bool)   //zero  channel


	go func() {
		for{
			// 空channel不能取出值，会阻塞
			// close(channel) 能一直取出值，但其more会为false
			j, more := <- jobs
			if more{
				fmt.Println("received job", j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <=3 ; j++{
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	//等待另外一个goroutine
	<- done
	fmt.Println("waiting  END")

}
