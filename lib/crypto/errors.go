package crypto

import "errors"

var (
	ErrEncrypt      = errors.New("encrypt Error:")
	ErrDecrypt      = errors.New("decrypt Error:")
	ErrTextIsShort  = errors.New("text/key too short or long!")
	ErrHashNotMatch = errors.New("hash not match")
)
