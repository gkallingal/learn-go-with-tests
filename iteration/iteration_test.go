package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("x", 10)
	want := "xxxxxxxxxx"

	if got != want {
		t.Errorf("Wanted %q, Got %q\n", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x", 10)
	}
}
