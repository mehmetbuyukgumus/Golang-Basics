package main

import (
	"fmt"
	"sync"
	"time"
)

// Exapmle 1
func func1() {
	fmt.Println("f1")
}

func func2() {
	fmt.Println("f2")
}

func main() {
	//Example 1
	go func1()
	go func2()
	time.Sleep(1 * time.Second)
	fmt.Println("End of the Example 1")

	//Example 2
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func ()  {
		defer wg.Done()
		fmt.Println("f1")
	}()
	go func ()  {
		defer wg.Done()
		fmt.Println("f2")
	}()
	wg.Wait()
	fmt.Println("End of the Example 2")
}
