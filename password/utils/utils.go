package utils

import (
	"fmt"
	"math/rand/v2"

	"github.com/fatih/color"
)

func GetUserInput(label string) string {
	var result string

	fmt.Printf(color.CyanString("%v: "), label)
	fmt.Scanln(&result)

	return result
}

func GeneratePassword(strLength int) string {
	runes := "ABCDEFJKLMNOPQRSTUVWXYZabcdefjklmnopqrstuvwxyz1234567890!@#$%^&*"
	result := ""

	for i := 0; i < strLength; i++ {
		randomRune := string(runes[rand.IntN(len(runes))])
		result = result + randomRune
	}

	return result
}

func RenderList(items []string) {
	fmt.Println("")
	for index, menuItem := range items {
		choiceNumber := fmt.Sprintf("[%v]", index+1)
		str := fmt.Sprintf("%v %v\n", color.YellowString(choiceNumber), menuItem)
		fmt.Print(str)
	}
	fmt.Println("")
}
