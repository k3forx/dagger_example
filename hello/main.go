package main

import (
	"fmt"
	"os"

	"github.com/k3forx/dagger_example/hello/greeting"
)

func main() {
	name := os.Getenv("NAME")
	if name == "" {
		name = "John Doe"
	}
	fmt.Printf(greeting.Greeting(name))
}
