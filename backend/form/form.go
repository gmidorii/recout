package form

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

const defaultLimit = 10

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
