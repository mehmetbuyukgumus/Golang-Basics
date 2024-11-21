package main

import (
	"fmt"
)

func main() {
	total := 0
	var names = [3]string{"John", "Jack", "Juliet"}

	for index := 0; index <= 10; index++ {
		total = index + total
	}
	fmt.Println(total)

	for index := 0; index < len(names); index++ {
		fmt.Println(names[index])
		if names[index] == "Jack" {
			break
		}
	}

	for index := 0; index <= 10; index++ {
		if index == 3 || index == 5 {
			fmt.Println(index)
		} else {
			continue
		}
	}
}
