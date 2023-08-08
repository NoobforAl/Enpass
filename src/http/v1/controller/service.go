package controller

import (
	"net/http"

	"github.com/NoobforAl/Enpass/contract"
	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func NewService(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser schema.Service
		var err error

		if err = validator.
			ParsService(c, &ser); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(ser, 0)
		service, err = interactor.
			New(stor, logger).
			CreateService(c, service)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func AllService(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		services, err := interactor.
			New(stor, logger).
			GetAllService(c)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

func FindService(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(
			schema.Service{}, uint(id),
		)

		service, err = interactor.
			New(stor, logger).
			FindService(c, service)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func UpdateService(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser schema.Service
		var err error

		if err = validator.
			ParsService(c, &ser); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(ser, uint(id))

		service, err = interactor.
			New(stor, logger).
			UpdateService(c, service)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func DeleteService(
	stor contract.Store,
	validator contract.Validation,
	logger contract.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(
			schema.Service{}, uint(id),
		)

		service, err = interactor.
			New(stor, logger).
			DeleteService(c, service)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}
