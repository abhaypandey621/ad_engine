package model

import "errors"

var (
	ErrInvalidAppIdentifier = errors.New("invalid app identifier")
	ErrBadRequest           = errors.New("bad request")
)
