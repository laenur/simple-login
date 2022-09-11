package constant

import "errors"

var (
	ErrInternal = errors.New("internal error")
	ErrNotFound = errors.New("not found")

	ErrInvalidUserID = errors.New("invalid UserID")
	ErrInvalidToken  = errors.New("invalid access token")
)
