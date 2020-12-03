package slices

import "testing"

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := SumSlice(numbers)

	want := 6

	if got != want {
		t.Errorf("Got %d, Want %d, Given %v\n", got, want, numbers)
	}
}
