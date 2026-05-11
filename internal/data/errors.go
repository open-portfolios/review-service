package data

import "errors"

var (
	ErrAlreadyReplied      = errors.New("comment already replied")
	ErrHorizontalOverreach = errors.New("horizontal overreach")
)
