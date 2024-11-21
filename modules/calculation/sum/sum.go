package sum

func Sum(numbers []int) int{
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}