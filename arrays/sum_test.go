package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("calculate sum of a slice", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)

		want := 12

		if got != want {
			t.Errorf("Got %d, Wanted %d, Given %v\n", got, want, numbers)
		}
	})

	t.Run("calculate multiple slice and return total into another slice", func(t *testing.T) {
		got := SumAll([]int{2, 4, 6}, []int{3, 4, 5})

		want := []int{12, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, Wanted %v\n", got, want)
		}
	})
}
