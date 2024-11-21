package main

import (
	"fmt"
)

func main() {
	var note int = 70

	if note >= 90 {
		fmt.Println("Your Grade: A")
	} else if note <= 80  && note >= 70{
		fmt.Println("Your Grade: B")
	} else {
		fmt.Println("Your Grade: C")
	}

	
}
