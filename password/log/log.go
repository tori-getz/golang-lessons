package log

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(message any) {
	str := fmt.Sprintf("%v", message)
	color.Green(str)
}

func Error(message any) {
	switch t := message.(type) {
	case string:
		str := fmt.Sprintf("%v", t)
		color.Red(str)
	case int:
		str := fmt.Sprintf("Error code: %v", t)
		color.Red(str)
	default:
		color.Red("Unknown error")
	}
}
