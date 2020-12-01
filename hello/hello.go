package main

import "fmt"

const (
	englishPrefix = "Hello, "
	spanish       = "Spanish"
	french        = "French"
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func main() {
	//	fmt.Println("Hello, World")
	fmt.Println(SayHello("Pippy", ""))
}

func SayHello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return prefixGreeting(language) + name

}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
