package main

const (
	recoutEntityName = "RecoutEntity"
	userEntityName   = "UserEntity"
)

type RecoutEntity struct {
	AccountID string
	Message   string
	CreatedAt int64
}

type UserEntity struct {
	AccountID string
	PixelaURL string
}
