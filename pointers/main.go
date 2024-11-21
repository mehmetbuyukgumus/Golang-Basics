package main

import "fmt"

func main()  {
	var number int = 12
	var ram_adress *int = &number
	*ram_adress = 20
	fmt.Println(number)

	var a int = 10
	var b int = 12
	var p *int = &a
	*p = b
	fmt.Println(a)
}