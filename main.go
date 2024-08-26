package main

import (
	"fmt"
	"github.com/JustTimmm/GoColor/color"
)

func main() {}

func colorLog(colorMessage string, message string, useBackgroundColor bool, backgroundColor string, args ...any) {
	if useBackgroundColor {
		if len(args) > 0 {
			fmt.Printf(backgroundColor+colorMessage+message+color.Reset, args...)
		} else {
			fmt.Printf(backgroundColor+colorMessage+message+color.Reset, args...)
		}
	} else {
		if len(args) > 0 {
			fmt.Printf(colorMessage+message+color.Reset, args...)
		} else {
			fmt.Printf(colorMessage+message+color.Reset, args...)
		}
	}
}

func errorLog(message string, args ...any) {
	colorLog(color.LightRed, message, false, "", args...)
}

func successLog(message string, args ...any) {
	colorLog(color.LightGreen, message, false, "", args...)
}

func infoLog(message string, args ...any) {
	colorLog(color.LightBlue, message, false, "", args...)
}

func debugLog(message string, args ...any) {
	colorLog(color.LightMagenta, message, false, "", args...)
}

func warnLog(message string, args ...any) {
	colorLog(color.LightYellow, message, false, "", args...)
}
