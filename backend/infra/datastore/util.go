package datastore

import (
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"go.mercari.io/datastore"
)

const (
	namespacePrefix = "recout"
)

func generateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(err, "failed generate uuid")
	}
	return id.String(), nil
}

func generateNamespace(env string) string {
	fmt.Printf("%v_%v\n", namespacePrefix, env)
	return fmt.Sprintf("%v_%v", namespacePrefix, env)
}

func newQuery(c datastore.Client, kind, env string) datastore.Query {
	return c.NewQuery(kind).Namespace(generateNamespace(env))
}

func generateKey(client datastore.Client, kind, env, uid string) datastore.Key {
	key := client.NameKey(kind, uid, nil)
	key.SetNamespace(generateNamespace(env))
	return key
}
