package main

import (
	"errors"
	"fmt"
)

const jsonTokenLeftBrace = "{"
const jsonTokenRightBrace = "}"

func lex(input string) ([]string, error) {
	var tokens []string

	for length := len(input); length > 0; {
		var c = string(input[0])

		if c == jsonTokenLeftBrace || c == jsonTokenRightBrace {
			tokens = append(tokens, c)

			input = input[1:]
		} else {
			fmt.Printf("Invalid character: %s\n", c)

			return tokens, errors.New("Invalid character")
		}
	}

	return tokens, nil
}

func parse(tokens []string) (interface{}, error) {
	return nil, nil
}

// FromJSON parses the JSON input and populates the output variable.
func FromJSON(input string, output interface{}) error {
	var tokens, err = lex(string(input))

	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		return errors.New("Empty input")
	}

	var result, err2 = parse(tokens)

	if err2 != nil {
		return err2
	}

	output = result

	return nil
}
