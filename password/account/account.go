package account

import (
	"errors"
	"net/url"
	"password/app/timestamp"
	"password/app/utils"
	"time"
)

const InvalidLogin = "INVALID_LOGIN"
const InvalidURL = "INVALID_URL"

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
	timestamp.Timestamp
}

func NewAccount(login string, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)

	if login == "" {
		return nil, errors.New(InvalidLogin)
	}

	if err != nil {
		return nil, errors.New(InvalidURL)
	}

	acc := &Account{
		Login:    login,
		Password: utils.GeneratePassword(12),
		Url:      urlString,
		Timestamp: timestamp.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return acc, nil
}
