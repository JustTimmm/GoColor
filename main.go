package main

import (
	"GoColor/color"
	"fmt"
)

func main() {}

func colorLog(colorMessage, message string, args ...any) {
	if len(args) > 0 {
		fmt.Printf(colorMessage+message+color.Reset, args...)
	} else {
		fmt.Print(colorMessage + message + color.Reset)
	}
}

func errorLog(message string, args ...any) {
	colorLog(color.Red, message, args...)
}

func successLog(message string, args ...any) {
	colorLog(color.Green, message, args...)
}

func infoLog(message string, args ...any) {
	colorLog(color.Blue, message, args...)
}

func debugLog(message string, args ...any) {
	colorLog(color.Magenta, message, args...)
}

func warnLog(message string, args ...any) {
	colorLog(color.Yellow, message, args...)
}
