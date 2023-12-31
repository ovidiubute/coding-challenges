package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func numberOfCharacters(file *os.File) int {
	var count int = 0
	var buffer []byte = make([]byte, 1024)

	for {
		n, err := file.Read(buffer)

		if err != nil {
			break
		}

		count += n
	}

	return count
}

func numberOfWords(file *os.File) int {
	var count int = 0
	var reader = bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count += len(strings.Fields(line))
	}

	return count
}

func numberOfLines(file *os.File) int {
	var count int = 0
	var reader = bufio.NewReader(file)

	for {
		_, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count++
	}

	return count
}

func main() {
	args := os.Args[1:] // Get the command line arguments excluding the program name

	if len(args) == 0 {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	// Read flags as the first argument
	flags := make(map[string]bool) // Create a map to store the flags

	// Read file name as the second argument
	filename := "" // Create a variable to store the file name

	// Loop through the arguments
	for _, arg := range args {
		// Check if the argument starts with a dash
		if arg[0] == '-' {
			// Loop through the flags
			for _, flag := range arg[1:] {
				flags[string(flag)] = true // Add the flag to the map
			}
		} else {
			filename = arg // Store the file name
		}
	}

	// Read file contents as a byte slice
	if filename != "" {
		file, err := os.Open(filename) // Open the file
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer file.Close() // Close the file after the function returns

		// Perform work depending on flags
		var base = path.Base(filename)
		if flags["l"] {
			fmt.Println(numberOfLines(file), " ", base)
		}

		if flags["c"] {
			fmt.Println(numberOfCharacters(file), " ", base)
		}

		if flags["w"] {
			fmt.Println(numberOfWords(file), " ", base)
		}
	}

	os.Exit(0)
}
