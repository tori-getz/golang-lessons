package main

import (
	"fmt"
	"password/app/account"
	"password/app/encrypter"
	"password/app/env"
	"password/app/files"
	"password/app/log"
	"password/app/menu"
	"password/app/utils"
	"strings"

	"github.com/fatih/color"
)

var menuMap = map[string]func(*account.VaultDecorated){
	menu.NEW_ACCOUNT:           newAccount,
	menu.FIND_ACCOUNT_BY_URL:   findAccountByUrl,
	menu.FIND_ACCOUNT_BY_LOGIN: findAccountByLogin,
	menu.REMOVE_ACCOUNT:        removeAccount,
	menu.EXIT:                  exit,
}

func main() {
	color.Magenta("__ password app __")

	env.LoadEnv()

	vault, err := account.NewVault(
		files.NewJsonDatabase("data.vault"),
		encrypter.NewEncrypter(),
	)

	if err != nil {
		panic(err.Error())
	}

Menu:
	for {
		choice := menu.ShowMenu()
		menuFunc := menuMap[choice]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(vault)
	}
}

func newAccount(vault *account.VaultDecorated) {
	login := utils.GetUserInput("Enter login")
	url := utils.GetUserInput("Enter URL")

	vault.AddAccount(login, url)
}

func findAccountByUrl(vault *account.VaultDecorated) {
	findUrl := utils.GetUserInput("Enter URL")
	findAccounts, err := vault.FindAccounts(func(acc *account.Account) bool {
		return strings.Contains(acc.Url, findUrl)
	})

	if err != nil {
		log.Error("Accounts not found!")
		return
	}

	renderAccountList(findAccounts)
}

func findAccountByLogin(vault *account.VaultDecorated) {
	findLogin := utils.GetUserInput("Enter login")
	findAccounts, err := vault.FindAccounts(func(acc *account.Account) bool {
		return strings.Contains(acc.Login, findLogin)
	})

	if err != nil {
		log.Error("Accounts not found!")
		return
	}

	renderAccountList(findAccounts)
}

func removeAccount(vault *account.VaultDecorated) {
	findUrl := utils.GetUserInput("Enter url")
	err := vault.RemoveAccount(findUrl)

	if err != nil {
		log.Error("Account not found!")
	}

	successStr := fmt.Sprintf("Account %v removed successfuly!", findUrl)
	log.Success(successStr)
}

func exit(_ *account.VaultDecorated) {
	log.Success("Have a nice day!")
	return
}

func renderAccountList(accounts []account.Account) {
	formattedAccounts := []string{}

	for _, account := range accounts {
		str := fmt.Sprintf("%v - Login: %v Password: %v", account.Url, account.Login, account.Password)
		formattedAccounts = append(formattedAccounts, str)
	}

	utils.RenderList(formattedAccounts)
}
