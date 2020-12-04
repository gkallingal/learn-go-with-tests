package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("calculate sum of an array", func(t *testing.T) {
		//declare array
		numbers := []int{1, 1, 1, 1, 2}

		got := Sum(numbers)

		want := 6

		if got != want {
			t.Errorf("Expected: %d, Got: %d, Given: %v\n", want, got, numbers)
		}
	})

	t.Run("calculate sum of a slice", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)

		want := 12

		if got != want {
			t.Errorf("Got %d, Wanted %d, Given %v\n", got, want, numbers)
		}
	})
}
