package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

func numberOfBytes(reader *bufio.Reader) int {
	var count int = 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count += len(line)
	}

	return count
}

func numberOfCharacters(reader *bufio.Reader) int {
	var count int = 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count += utf8.RuneCountInString(line)
	}

	return count
}

func numberOfWords(reader *bufio.Reader) int {
	var count int = 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count += len(strings.Fields(line))
	}

	return count
}

func numberOfLines(reader *bufio.Reader) int {
	var count int = 0

	for {
		_, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		count++
	}

	return count
}

func printResults(filename string, result int) {
	var displayFilename string = filename
	if filename != "" {
		displayFilename = path.Base(filename)
	}

	fmt.Println(result, " ", displayFilename)

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
	filename := ""

	// Loop through the arguments
	for _, arg := range args {
		// Check if the argument starts with a dash
		if arg[0] == '-' {
			// Loop through the flags
			for _, flag := range arg[1:] {
				flags[string(flag)] = true
			}
		} else {
			filename = arg // Store the file name
		}
	}

	var reader *bufio.Reader
	if filename == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(filename)

		if err != nil {
			fmt.Println("Error opening file")
			os.Exit(1)
		}

		defer file.Close()

		reader = bufio.NewReader(file)
	}

	// Perform work depending on flags

	if flags["l"] {
		var lines int = numberOfLines(reader)
		printResults(filename, lines)
	}

	if flags["m"] {
		var characters = numberOfCharacters(reader)

		printResults(filename, characters)
	}

	if flags["w"] {
		var words int = numberOfWords(reader)
		printResults(filename, words)
	}

	if flags["c"] {
		var bytes int = numberOfBytes(reader)

		printResults(filename, bytes)
	}

	os.Exit(0)
}
