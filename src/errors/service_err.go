package errors

import "errors"

var (
	ErrNotFoundUser   = errors.New("Not found user")
	ErrUpdatePassDone = errors.New("update password with all password done but")
)
