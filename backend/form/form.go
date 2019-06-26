package form

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

const defaultLimit = 10

type Recout struct {
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

type RecoutFetch struct {
	AccountID string
	Limit     int
}

func FactoryFetchForm(values url.Values) (RecoutFetch, error) {
	accountID := values.Get("account_id")
	if accountID == "" {
		return RecoutFetch{}, xerrors.New("required 'account_id' param")
	}
	limit, err := getIntValue(values, "limit", defaultLimit)
	if err != nil {
		return RecoutFetch{}, errors.Wrap(err, "failed parse query key=limit")
	}
	return RecoutFetch{AccountID: accountID, Limit: limit}, nil
}

type RecoutContinues struct {
	AccountID string
}

func FactoryContinues(values url.Values) (RecoutContinues, error) {
	id := values.Get("account_id")
	if id == "" {
		return RecoutContinues{}, errors.New("account_id is necessary parameter")
	}
	return RecoutContinues{AccountID: id}, nil
}

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

func getIntValue(values url.Values, key string, defaultValue int) (int, error) {
	v := values.Get(key)
	if v == "" {
		return defaultValue, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.Wrap(err, "failed a to i")
	}

	return i, nil
}
