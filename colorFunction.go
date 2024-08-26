package GoColor

import "fmt"

func colorLog(colorMessage string, message string, useBackgroundColor bool, backgroundColor string, args ...any) {
	if useBackgroundColor {
		if len(args) > 0 {
			fmt.Printf(backgroundColor+colorMessage+message+Reset, args...)
		} else {
			fmt.Printf(backgroundColor+colorMessage+message+Reset, args...)
		}
	} else {
		if len(args) > 0 {
			fmt.Printf(colorMessage+message+Reset, args...)
		} else {
			fmt.Printf(colorMessage+message+Reset, args...)
		}
	}
}

func errorLog(message string, args ...any) {
	colorLog(LightRed, message, false, "", args...)
}

func successLog(message string, args ...any) {
	colorLog(LightGreen, message, false, "", args...)
}

func infoLog(message string, args ...any) {
	colorLog(LightBlue, message, false, "", args...)
}

func debugLog(message string, args ...any) {
	colorLog(LightMagenta, message, false, "", args...)
}

func warnLog(message string, args ...any) {
	colorLog(LightYellow, message, false, "", args...)
}
