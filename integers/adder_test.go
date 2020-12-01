package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("Expected %d but received %d\n", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)

	//Output: 5
}
