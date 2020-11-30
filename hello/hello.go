package main

const (
	englishPrefix = "Hello, "
)

func main() {
	//	fmt.Println("Hello, World")
	SayHello("Pippy")
}

func SayHello(name string) string {
	return englishPrefix + name
}
