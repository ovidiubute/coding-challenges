package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get the command line arguments excluding the program name

	// Print the arguments
	for i, arg := range args {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}
}
