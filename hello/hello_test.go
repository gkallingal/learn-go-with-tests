package main

import "testing"

func TestHelloWorld(t *testing.T) {
	request := SayHello("George")
	response := "Hello, George"

	if request != response {
		t.Errorf("There was an error.  Expected to get %q but received %q\n", response, request)
	}
}
