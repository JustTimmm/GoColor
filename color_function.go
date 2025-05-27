package GoColor

import (
	"fmt"
	"strings"
)

type ColorOption struct {
	TextColor       string
	BackgroundColor string
}

func ColorLog(option ColorOption, message string, args ...any) {
	formatted := fmt.Sprintf(message, args...)
	lines := strings.Split(formatted, "\n")
	for i, line := range lines {
		if i < len(lines)-1 {
			fmt.Print(option.BackgroundColor + option.TextColor + line + Reset + "\n")
		} else {
			fmt.Print(option.BackgroundColor + option.TextColor + line + Reset)
		}
	}
}

func ErrorLog(message string, args ...interface{}) {
	ColorLog(ColorOption{TextColor: LightRed}, message, args...)
}

func SuccessLog(message string, args ...interface{}) {
	ColorLog(ColorOption{TextColor: LightGreen}, message, args...)
}

func InfoLog(message string, args ...interface{}) {
	ColorLog(ColorOption{TextColor: LightBlue}, message, args...)
}

func DebugLog(message string, args ...interface{}) {
	ColorLog(ColorOption{TextColor: LightMagenta}, message, args...)
}

func WarnLog(message string, args ...interface{}) {
	ColorLog(ColorOption{TextColor: LightYellow}, message, args...)
}

func RainbowLog(message string, args ...interface{}) {
	colors := []string{Red, Yellow, Green, Cyan, Blue, Magenta}
	colorCount := len(colors)

	formatted := fmt.Sprintf(message, args...)

	for i, char := range formatted {
		color := colors[i%colorCount]
		fmt.Print(color + string(char) + Reset)
	}
}
