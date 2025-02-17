package menu

import "password/app/utils"

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

	utils.RenderList(menuItems[:])

	enteredChoice := utils.GetUserInput("Enter menu number")

	return enteredChoice
}
