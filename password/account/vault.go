package account

import (
	"encoding/json"
	"password/app/files"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() (*Vault, error) {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}, nil
}

func (vault *Vault) AddAccount(login string, url string) (*Account, error) {
	newAcc, err := New(login, url)

	if err != nil {
		return nil, err
	}

	vault.Accounts = append(vault.Accounts, *newAcc)
	vault.UpdatedAt = time.Now()

	bytes, err := vault.ToBytes()

	if err != nil {
		return nil, err
	}

	files.Write("vault.json", bytes)

	return newAcc, nil
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
