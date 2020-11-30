package main

const (
	englishPrefix = "Hello, "
)

func main() {
	//	fmt.Println("Hello, World")
	SayHello("Pippy")
}

func SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
