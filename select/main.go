package main

import (
	"fmt"
	"time"
)


func main()  {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(10* time.Second)
		channel1 <- "Hello"
	}()

	go func() {
		time.Sleep(1* time.Second)
		channel2 <- "World"
	}()

	var data1 string = <- channel1
	var data2 string = <- channel2

	fmt.Println("Data 1: ", data1)
	fmt.Println("Data 2: ", data2)


}