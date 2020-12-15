package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GreetName(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	GreetName(os.Stdout, "George")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler))
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	GreetName(w, "George")
}
