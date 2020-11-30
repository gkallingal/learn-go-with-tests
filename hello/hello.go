package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spanish       = "Spanish"
	spanishPrefix = "Hola, "
)

func main() {
	//	fmt.Println("Hello, World")
	fmt.Println(SayHello("Pippy", ""))
}

func SayHello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishPrefix + name
	default:
		return englishPrefix + name
	}

}
