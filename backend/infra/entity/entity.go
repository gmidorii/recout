package entity

const (
	RecoutEntityName = "RecoutEntity"
	UserEntityName   = "UserEntity"
)

type Recout struct {
	AccountID string `datastore:"account_id"`
	Message   string `datastore:"message,noindex"`
	CreatedAt int64  `datastore:"created_at"`
}

type User struct {
	AccountID   string `datastore:"account_id"`
	AccessToken string `datastore:"access_token"`
	Name        string `datastore:"name"`
	PixelaGraph string `datastore:"pixela_graph"`
	PixelaToken string `datastore:"pixela_token"`
}
