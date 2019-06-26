package form

import (
	"net/url"

	"github.com/pkg/errors"
)

type User struct {
	AccountID   string `json:"account_id"`
	PixelaGraph string `json:"pixela_graph"`
	PixelaToken string `json:"pixela_token"`
}

func FactoryUser(values url.Values) (User, error) {
	id := values.Get("account_id")
	if id == "" {
		return User{}, errors.New("account_id is necessary paramter")
	}
	return User{AccountID: id}, nil
}
