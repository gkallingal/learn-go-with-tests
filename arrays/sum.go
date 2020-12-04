package arrays

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAllTails(numbers ...[]int) (sum []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(number[1:]))
		}
	}
	return
}
