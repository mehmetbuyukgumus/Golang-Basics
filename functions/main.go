package main

import "fmt"

var names = []string{"Python", "Java", "Golang", "C++"}

func main() {
	result := add(10, 20)
	fmt.Println(result)
	say_hi("Mehmet")
	fmt.Printf("Our Technologies:\n")
	techs()
	result_1, result_2 := calculation(10, 8)
	fmt.Println(result_1, result_2)
	result_sum := sum([]int{2,3,7,4,2,65,12})
	fmt.Println(result_sum)
	fmt.Println(sum_all_numbers(1,2,3,4,5))

}

func add(x int, y int) int {
	var sonuc int = x + y
	return sonuc
}

func say_hi(name string) {
	fmt.Printf("Hi %s\n", name)
}

func techs() {
	for _, value := range names {
		fmt.Printf("%s\n", value)
	}
}

func calculation(x int, y int) (int, int) {
	minus := x - y
	plus := x + y
	return minus, plus
}

func sum (numbers []int)int {
	total := 0
	for _, value := range(numbers){
		total += value
	}
	return total
}

func sum_all_numbers(number ...int)int {
	total := 0
	for _, value := range(number){
		total += value
	}
	return total
}