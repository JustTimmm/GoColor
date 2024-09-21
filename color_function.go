package GoColor

import "fmt"

type ColorOption struct {
	TextColor       string
	BackgroundColor string
}

func ColorLog(option ColorOption, message string, args ...any) {
	format := option.BackgroundColor + option.TextColor
	msg := fmt.Sprintf(format+message+Reset, args...)

	fmt.Printf(msg + Reset)
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
