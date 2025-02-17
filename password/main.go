package main

import (
	"password/app/account"
	"password/app/menu"
	"password/app/utils"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("__ password app __")

	vault, err := account.NewVault()

	if err != nil {
		panic(err)
	}

Menu:
	for {
		choice := menu.ShowMenu()

		switch choice {
		case menu.NEW_ACCOUNT:
			NewAccount(vault)

		case menu.FIND_ACCOUNT:
			FindAccount()

		case menu.REMOVE_ACCOUNT:
			RemoveAccount()

		case menu.EXIT:
			color.Green("Have a nice day!")
			break Menu
		}
	}
}

func NewAccount(vault *account.Vault) {
	login := utils.GetUserInput("Enter login")
	url := utils.GetUserInput("Enter URL")

	vault.AddAccount(login, url)
}

func FindAccount() {
	panic("to do")
}

func RemoveAccount() {
	panic("to do")
}
