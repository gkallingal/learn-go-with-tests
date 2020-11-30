package main

import "fmt"

const (
	englishPrefix = "Hello, "
)

func main() {
	//	fmt.Println("Hello, World")
	fmt.Println(SayHello("Pippy", ""))
}

func SayHello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
