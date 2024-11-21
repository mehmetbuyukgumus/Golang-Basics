package main

import (
	"fmt"
)

func main()  {
	names:= map[string]int{
		"Paris": 0,
		"NYC": 1,
		"London":2,
	}
	
	delete(names, "NYC")
	fmt.Println(names)
}