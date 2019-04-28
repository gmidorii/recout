package main

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

const defaultLimit = 10

type RecoutForm struct {
	Message string `json:message`
}

type FetchForm struct {
	Limit int
}

func FactoryFetchForm(values url.Values) (FetchForm, error) {
	limit, err := getIntValue(values, "limit", defaultLimit)
	if err != nil {
		return FetchForm{}, errors.Wrap(err, "failed parse query key=limit")
	}
	return FetchForm{Limit: limit}, nil
}

type UserForm struct {
	AccountID string `json:"account_id"`
	Graph     string `json:"graph"`
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
