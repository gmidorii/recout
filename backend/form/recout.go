package form

import (
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

type Recout struct {
	AccountID string `json:"account_id"`
	Message   string `json:"message"`
}

type RecoutFetch struct {
	AccountID string
	Limit     int
}

type RecoutContinues struct {
	AccountID string
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

func FactoryContinues(values url.Values) (RecoutContinues, error) {
	id := values.Get("account_id")
	if id == "" {
		return RecoutContinues{}, errors.New("account_id is necessary parameter")
	}
	return RecoutContinues{AccountID: id}, nil
}
