package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInputFromTerminal(prompt string) (string, error) {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var text string
	var err error

	// Optionally print input prompt
	if prompt != "" {
		fmt.Printf("%s: ", prompt)
	}

	// Read an input from the terminal
	text, err = reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error reading input: %s", err.Error())
	}

	// Remove trailing whitespace
	text = strings.TrimSpace(text)

	return text, nil
}

func ReadIntegerFromTerminal(prompt string) (int, error) {
	var text string
	var err error
	var number int

	// Get input from terminal
	if text, err = ReadInputFromTerminal(prompt); err != nil {
		return 0, err
	}

	// Convert the input to an integer
	if number, err = strconv.Atoi(text); err != nil {
		return 0, fmt.Errorf("Error converting input to 'int': %s", err.Error())
	}

	return number, nil
}

func ReadFloatFromTerminal(prompt string) (float64, error) {
	var text string
	var err error
	var number float64

	// Get input from terminal
	if text, err = ReadInputFromTerminal(prompt); err != nil {
		return 0, err
	}

	// Convert the input to an integer
	if number, err = strconv.ParseFloat(text, 64); err != nil {
		return 0, fmt.Errorf("Error converting input to 'float64': %s", err.Error())
	}

	return number, nil
}

func ReadBooleanFromTerminal(prompt string, trueVals, falseVals []string, caseSensitive bool) (bool, error) {
	var text string
	var err error

	// Get input from terminal
	if text, err = ReadInputFromTerminal(prompt); err != nil {
		return false, err
	}

	// Optionally, apply case-insensitivity
	if !caseSensitive {
		text = strings.ToLower(text)
		for i:=0; i<len(trueVals); i++ {
			trueVals[i] = strings.ToLower(trueVals[i])
		}
		for i:=0; i<len(falseVals); i++ {
			falseVals[i] = strings.ToLower(falseVals[i])
		}
	}

	// Convert the input to a boolean
	if StringArrayContains(trueVals, text) {
		return true, nil
	} else if StringArrayContains(falseVals, text) {
		return false, nil
	}

	return false, fmt.Errorf("Invalid input '%s'", text)
}
