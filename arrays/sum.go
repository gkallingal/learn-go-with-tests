package arrays

func Sum(numbers [5]int) int {
	var result int
	for i := 0; i < len(numbers); i++ {
		result = result + numbers[i]
	}
	return result
}
