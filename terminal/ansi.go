package terminal

import (
	"fmt"
)

type Colour int

const (
	Black = 0
	White = 15

	Blue      = 21
	LightBlue = 81
	DarkBlue  = 18

	Green = 34

	Orange = 214

	Red     = 160
	DarkRed = 88
)

/*****************************************************************************
	Cursor manipulation
*****************************************************************************/

func CursorUp(lines int) {
	fmt.Printf("\033[%dA", lines)
}

func CursorDown(lines int) {
	fmt.Printf("\033[%dB", lines)
}

func CursorRight(columns int) {
	fmt.Printf("\033[%dC", columns)
}

func CursorLeft(columns int) {
	fmt.Printf("\033[%dD", columns)
}

func SetCursorColumn(column int) {
	fmt.Printf("\033[%dG", column)
}

func SetCursorPosition(line, column int) {
	fmt.Printf("\033[%d;%dH", line, column)
}

func SetCursorHome() {
	fmt.Print("\033[H")
}

/*****************************************************************************
	Erasing information from screen
*****************************************************************************/

// Special case of EraseScreen() that also resets the cursor to 0,0
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func EraseScreen() {
	fmt.Print("\033[2J")
}

func EraseCurrentLine() {
	fmt.Print("\033[2K")
}

func EraseFromStartOfLine() {
	fmt.Print("\033[1K")
}

func EraseFromStartOfScreen() {
	fmt.Print("\033[1J")
}

func EraseToEndOfLine() {
	fmt.Print("\033[0K")
}

func EraseToEndOfScreen() {
	fmt.Print("\033[0J")
}

/*****************************************************************************
	Set font styles and colours
*****************************************************************************/

func SetBold(enable bool) {
	if enable {
		fmt.Print("\033[1m")
	} else {
		fmt.Print("\033[22m")
	}
}

func SetFaint(enable bool) {
	if enable {
		fmt.Print("\033[2m")
	} else {
		fmt.Print("\033[22m")
	}
}

func SetItalic(enable bool) {
	if enable {
		fmt.Print("\033[3m")
	} else {
		fmt.Print("\033[23m")
	}
}

func SetUnderline(enable bool) {
	if enable {
		fmt.Print("\033[4m")
	} else {
		fmt.Print("\033[24m")
	}
}

func SetBlink(enable bool) {
	if enable {
		fmt.Print("\033[5m")
	} else {
		fmt.Print("\033[25m")
	}
}

func SetInverse(enable bool) {
	if enable {
		fmt.Print("\033[7m")
	} else {
		fmt.Print("\033[27m")
	}
}

func SetHidden(enable bool) {
	if enable {
		fmt.Print("\033[8m")
	} else {
		fmt.Print("\033[28m")
	}
}

func SetStrikethrough(enable bool) {
	if enable {
		fmt.Print("\033[9m")
	} else {
		fmt.Print("\033[29m")
	}
}

func SetFontColour(colour Colour) {
	fmt.Printf("\033[38;5;%dm", colour)
}

func SetBackgroundColour(colour Colour) {
	fmt.Printf("\033[48;5;%dm", colour)
}

func ResetStyle() {
	fmt.Print("\033[0m")
}
