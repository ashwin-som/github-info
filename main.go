package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args

	if len(args) < 2 {
        fmt.Println("Usage: go run main.go <your-name>")
        return
    }

	name := args[1]
	fmt.Printf("Welcome to the other side %s :)\n", name)
}
