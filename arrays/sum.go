package arrays

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbers ...[]int) (sum []int) {
	for _, number := range numbers {
		sum = append(sum, Sum(number))
	}
	return
}
