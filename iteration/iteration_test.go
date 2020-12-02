package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("x")
	want := "xxxxx"

	if got != want {
		t.Errorf("Wanted %q, Got %q\n", want, got)
	}
}
