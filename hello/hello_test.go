package main

import "testing"

func TestHelloWorld(t *testing.T) {

	validateMessage := func(t *testing.T, request, response string) {
		t.Helper()
		if request != response {
			t.Errorf("There was an error.  Expected to get %q but received %q\n", response, request)
		}
	}

	t.Run("Check for Hello, World", func(t *testing.T) {
		request := SayHello("", "")
		response := "Hello, World"

		validateMessage(t, request, response)
	})

	t.Run("Check for Hello, [Name]", func(t *testing.T) {
		request := SayHello("George", "")
		response := "Hello, George"

		validateMessage(t, request, response)
	})

	t.Run("Check for Spanish", func(t *testing.T) {
		request := SayHello("Javier", "Spanish")
		response := "Hola, Javier"

		validateMessage(t, request, response)
	})

	t.Run("Check for French", func(t *testing.T) {
		request := SayHello("Serge", "French")
		response := "Bonjour, Serge"

		validateMessage(t, request, response)
	})
}
