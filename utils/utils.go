package utils

import (
	"fmt"
	"os"
)

func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition == true {
		return trueValue
	}
	return falseValue
}

func ExitWithMessage(message string, params ...interface{}) {
	fmt.Printf(message, params...)
	os.Exit(0)
}

func ExitWithError(message string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, message, params...)
	os.Exit(1)
}
