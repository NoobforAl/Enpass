package model

import "errors"

var (
	EncryptError     = errors.New("encrypt Error:")
	DecryptError     = errors.New("decrypt Error:")
	ErrHashNotMatch  = errors.New("hash not match")
	TextIsShortError = errors.New("text too short")
)
