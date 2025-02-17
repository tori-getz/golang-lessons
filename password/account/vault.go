package account

import (
	"encoding/json"
	"errors"
	"password/app/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDatabase struct {
	Vault
	db files.JsonDatabase
}

func NewVault(db files.JsonDatabase) (*VaultWithDatabase, error) {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDatabase{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}, nil
	}

	var vault Vault

	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("vault.json is invalid, use empty vault")
		return &VaultWithDatabase{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}, nil
	}

	return &VaultWithDatabase{
		Vault: vault,
		db:    db,
	}, nil
}

func (vault *VaultWithDatabase) AddAccount(login string, url string) (*Account, error) {
	newAcc, err := NewAccount(login, url)

	if err != nil {
		return nil, err
	}

	vault.Accounts = append(vault.Accounts, *newAcc)

	vault.save()

	return newAcc, nil
}

func (vault *VaultWithDatabase) FindAccounts(findUrl string) ([]Account, error) {
	findAccounts := []Account{}

	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, findUrl) {
			continue
		}

		findAccounts = append(findAccounts, account)
	}

	if len(findAccounts) == 0 {
		return nil, errors.New("ACCOUNTS_NOT_FOUND")
	}

	return findAccounts, nil
}

func (vault *VaultWithDatabase) RemoveAccount(findUrl string) error {
	nextAccounts := []Account{}
	accountFinded := false

	for _, account := range vault.Accounts {
		if account.Url == findUrl {
			accountFinded = true
			continue
		}

		nextAccounts = append(nextAccounts, account)
	}

	if accountFinded == false {
		return errors.New("ACCOUNT_NOT_FOUND")
	}

	vault.Accounts = nextAccounts

	err := vault.save()

	if err != nil {
		return err
	}

	return nil
}

func (vault *VaultWithDatabase) save() error {
	vault.Vault.UpdatedAt = time.Now()

	bytes, err := vault.Vault.ToBytes()

	if err != nil {
		return err
	}

	vault.db.Write(bytes)

	return nil
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
