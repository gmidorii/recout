package entity

const (
	RecoutEntityName = "RecoutEntity"
	UserEntityName   = "UserEntity"
)

type Recout struct {
	AccountID string
	Message   string
	CreatedAt int64
}

type User struct {
	AccountID string
	PixelaURL string
}
