package main

import (
	"fmt"
	"password/app/account"
	"password/app/files"
	"password/app/menu"
	"password/app/utils"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("__ password app __")

	db := files.NewJsonDatabase("vault.json")
	vault, err := account.NewVault(*db)

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
			FindAccount(vault)

		case menu.REMOVE_ACCOUNT:
			RemoveAccount(vault)

		case menu.EXIT:
			color.Green("Have a nice day!")
			break Menu
		}
	}
}

func NewAccount(vault *account.VaultWithDatabase) {
	login := utils.GetUserInput("Enter login")
	url := utils.GetUserInput("Enter URL")

	vault.AddAccount(login, url)
}

func FindAccount(vault *account.VaultWithDatabase) {
	findUrl := utils.GetUserInput("Enter URL")
	findAccounts, err := vault.FindAccounts(findUrl)

	if err != nil {
		color.Red("Accounts not found!")
		return
	}

	formattedAccounts := []string{}

	for _, account := range findAccounts {
		str := fmt.Sprintf("%v - Login: %v Password: %v", account.Url, account.Login, account.Password)
		formattedAccounts = append(formattedAccounts, str)
	}

	utils.RenderList(formattedAccounts)
}

func RemoveAccount(vault *account.VaultWithDatabase) {
	findUrl := utils.GetUserInput("Enter url")
	err := vault.RemoveAccount(findUrl)

	if err != nil {
		color.Green("Account not found!")
	}

	successStr := fmt.Sprintf("Account %v removed successfuly!", findUrl)
	color.Green(successStr)
}
