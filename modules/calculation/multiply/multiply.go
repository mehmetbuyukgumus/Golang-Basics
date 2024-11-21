package multiply

func Multiply(numbers []int) int {
	total := 1
	for _, num := range numbers {
		total *= num
	}
	return total
}