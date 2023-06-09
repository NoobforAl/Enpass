package model

import "errors"

var (
	ErrEncrypt      = errors.New("encrypt Error:")
	ErrDecrypt      = errors.New("decrypt Error:")
	ErrHashNotMatch = errors.New("hash not match")
	ErrTextIsShort  = errors.New("text too short")
)
