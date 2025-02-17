package menu

import "password/app/utils"

const NEW_ACCOUNT = "1"
const FIND_ACCOUNT_BY_URL = "2"
const FIND_ACCOUNT_BY_LOGIN = "3"
const REMOVE_ACCOUNT = "4"
const EXIT = "5"

type choice = string

func ShowMenu() choice {
	menuItems := [5]string{
		"New account",
		"Find account by URL",
		"Find account by login",
		"Remove account",
		"Exit",
	}

	utils.RenderList(menuItems[:])

	enteredChoice := utils.GetUserInput("Enter menu number")

	return enteredChoice
}
