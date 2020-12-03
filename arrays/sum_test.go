package arrays

import "testing"

func TestSum(t *testing.T) {
	//declare array
	numbers := [5]int{1, 1, 1, 1, 2}

	got := Sum(numbers)

	want := 6

	if got != want {
		t.Errorf("Expected: %d, Got: %d, Given: %v\n", want, got, numbers)
	}
}
