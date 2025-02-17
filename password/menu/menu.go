package menu

import (
	"fmt"
	"password/app/utils"

	"github.com/fatih/color"
)

const NEW_ACCOUNT = "1"
const FIND_ACCOUNT = "2"
const REMOVE_ACCOUNT = "3"
const EXIT = "4"

type choice = string

func ShowMenu() choice {
	menuItems := [4]string{
		"New account",
		"Find account",
		"Remove account",
		"Exit",
	}

	fmt.Println("")
	for index, menuItem := range menuItems {
		choiceNumber := fmt.Sprintf("[%v]", index+1)
		str := fmt.Sprintf("%v %v\n", color.YellowString(choiceNumber), menuItem)
		fmt.Print(str)
	}
	fmt.Println("")

	enteredChoice := utils.GetUserInput("Enter menu number")

	return enteredChoice
}
