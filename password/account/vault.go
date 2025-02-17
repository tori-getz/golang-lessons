package account

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultDb interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type VaultEncrypter interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type VaultDecorated struct {
	Vault
	db  VaultDb
	enc VaultEncrypter
}

func NewVault(db VaultDb, enc VaultEncrypter) (*VaultDecorated, error) {
	file, err := db.Read()
	if err != nil {
		return &VaultDecorated{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}, nil
	}

	decryptedFile := enc.Decrypt(file)

	var vault Vault

	err = json.Unmarshal(decryptedFile, &vault)

	if err != nil {
		color.Red("Vault JSON is invalid, use empty vault")
		return &VaultDecorated{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}, nil
	}

	return &VaultDecorated{
		Vault: vault,
		db:    db,
		enc:   enc,
	}, nil
}

func (vault *VaultDecorated) AddAccount(login string, url string) (*Account, error) {
	newAcc, err := NewAccount(login, url)

	if err != nil {
		return nil, err
	}

	vault.Accounts = append(vault.Accounts, *newAcc)

	vault.save()

	return newAcc, nil
}

type accountValidator = func(*Account) bool

func (vault *VaultDecorated) FindAccounts(validator accountValidator) ([]Account, error) {
	findAccounts := []Account{}

	for _, account := range vault.Accounts {
		if validator(&account) {
			findAccounts = append(findAccounts, account)
		}
	}

	if len(findAccounts) == 0 {
		return nil, errors.New("ACCOUNTS_NOT_FOUND")
	}

	return findAccounts, nil
}

func (vault *VaultDecorated) RemoveAccount(findUrl string) error {
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

func (vault *VaultDecorated) save() error {
	vault.Vault.UpdatedAt = time.Now()

	bytes, err := vault.Vault.ToBytes()

	if err != nil {
		return err
	}

	encryptedData := vault.enc.Encrypt(bytes)

	vault.db.Write(encryptedData)

	return nil
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
