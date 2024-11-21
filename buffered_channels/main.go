package main

import "fmt"

func main() {
	bufferedChannel := make(chan int,4)

	go func() {
		for i:=0; i<=10; i++ {
			bufferedChannel <- i
			fmt.Println("Sent data: ", i)
		}
		close(bufferedChannel)
	}()

	for data:= range bufferedChannel {
		fmt.Println("Received Data: ", data)
	}
}
