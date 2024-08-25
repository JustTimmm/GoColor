package main

import (
	"GoColor/color"
	"fmt"
)

func main() {}

func colorLog(colorMessage, message string, args ...any) {
	fmt.Printf(colorMessage+message+color.Reset, args...)
}
