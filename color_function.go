package GoColor

import "fmt"

func ColorLog(colorMessage string, message string, useBackgroundColor bool, backgroundColor string, args ...any) {
	if useBackgroundColor {
		if len(args) > 0 {
			fmt.Printf(backgroundColor+colorMessage+message+Reset, args...)
			return
		}
		fmt.Printf(backgroundColor + colorMessage + message + Reset)
		return
	}

	if len(args) > 0 {
		fmt.Printf(colorMessage+message+Reset, args...)
		return
	}
	fmt.Printf(colorMessage + message + Reset)
	return
}

func ErrorLog(message string, args ...interface{}) {
	ColorLog(LightRed, message, false, "", args...)
}

func SuccessLog(message string, args ...interface{}) {
	ColorLog(LightGreen, message, false, "", args...)
}

func InfoLog(message string, args ...interface{}) {
	ColorLog(LightBlue, message, false, "", args...)
}

func DebugLog(message string, args ...interface{}) {
	ColorLog(LightMagenta, message, false, "", args...)
}

func WarnLog(message string, args ...interface{}) {
	ColorLog(LightYellow, message, false, "", args...)
}
