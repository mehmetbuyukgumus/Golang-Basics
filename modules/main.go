package main

import (
	"fmt"
	"modules/calculation/multiply"
	"modules/calculation/sum"
)

func main() {
	var numbers = []int{1,2,3,4}
	result_sum := sum.Sum(numbers)
	result_multiply := multiply.Multiply(numbers)
	fmt.Println(result_sum)
	fmt.Println(result_multiply)
}