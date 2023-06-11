package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/NoobforAl/Enpass/database"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFoundPass      = errors.New("Not found Password saved")
	ErrNotMatchPassword  = errors.New("not match password")
	ErrUnexpectedSigning = errors.New("unexpected signing method:")
)

func errorHandling(c *gin.Context, err error) {
	var status = http.StatusInternalServerError

	switch {
	case errors.Is(err, database.ErrRecordNotFound):
		status = http.StatusNotFound

	case errors.Is(err, strconv.ErrSyntax):
		status = http.StatusBadRequest

	case errors.Is(err, database.ErrDecrypt):
	case errors.Is(err, database.ErrEncrypt):
	case errors.Is(err, database.ErrTextIsShort):
		status = http.StatusUnauthorized
	}

	c.JSON(status, gin.H{
		"detail": err.Error(),
	})
}
