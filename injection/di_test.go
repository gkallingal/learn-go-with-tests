package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	GreetName(buffer, "George")

	got := buffer.String()
	want := "Hello, George"
	if got != want {
		t.Errorf("Got %q, Want %q", got, want)
	}
}
