package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestWC(t *testing.T) {
	// Create a test input string
	input := "Hello, world!\nThis is a test.\n"

	// Create a new reader with the test input string
	reader := bufio.NewReader(strings.NewReader(input))

	// Call the wc function and get the results
	var bytes = numberOfBytes(reader)

	reader.Reset(strings.NewReader(input))

	var characters = numberOfCharacters(reader)

	reader.Reset(strings.NewReader(input))

	words := numberOfWords(reader)

	reader.Reset(strings.NewReader(input))
	lines := numberOfLines(reader)

	// Check the expected results
	expectedBytes := len(input)
	if bytes != expectedBytes {
		t.Errorf("Expected number of bytes: %d, got: %d", expectedBytes, bytes)
	}

	expectedCharacters := len(input)
	if characters != expectedCharacters {
		t.Errorf("Expected number of characters: %d, got: %d", expectedCharacters, characters)
	}

	expectedWords := 6
	if words != expectedWords {
		t.Errorf("Expected number of words: %d, got: %d", expectedWords, words)
	}

	expectedLines := 2
	if lines != expectedLines {
		t.Errorf("Expected number of lines: %d, got: %d", expectedLines, lines)
	}
}
