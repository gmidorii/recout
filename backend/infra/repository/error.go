package repository

// NotFoundError is entity not found.
// But query is suceeded.
type NotFoundError struct{}

func (n NotFoundError) Error() string {
	return "not found entity."
}
