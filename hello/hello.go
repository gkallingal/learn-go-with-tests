package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	SayHello("Pippy")
}

func SayHello(name string) string {
	return "Hello, " + name
}
