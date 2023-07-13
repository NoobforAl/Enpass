package errors

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotMatchPassword  = errors.New("not match password")
	ErrUnexpectedSigning = errors.New("unexpected signing method:")
)

func ErrHandle(c *gin.Context, err error) {
	var status = http.StatusInternalServerError

	switch {
	case errors.Is(err, ErrRecordNotFound):
		status = http.StatusNotFound

	case errors.Is(err, strconv.ErrSyntax):
		status = http.StatusBadRequest

	case errors.Is(err, ErrDecrypt) ||
		errors.Is(err, ErrEncrypt) ||
		errors.Is(err, ErrTextIsShort) ||
		errors.Is(err, ErrNotMatchPassword):
		status = http.StatusUnauthorized
	}

	c.JSON(status, gin.H{
		"detail": err.Error(),
	})
}
