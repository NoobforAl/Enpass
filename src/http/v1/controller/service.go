package controller

import (
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"
	"github.com/NoobforAl/Enpass/http/v1/parser"
	"github.com/NoobforAl/Enpass/interactor"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func NewService(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser schema.Service
		err := conf.Validation.ParsService(c, &ser)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(ser, 0)
		service, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).CreateService(c, service)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func AllService(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		services, err := interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).GetAllService(c)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

func FindService(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(
			schema.Service{}, uint(id),
		)

		service, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).FindService(c, service)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func UpdateService(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser schema.Service
		err := conf.Validation.ParsService(c, &ser)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(ser, uint(id))

		service, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).UpdateService(c, service)

		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

func DeleteService(conf *BaseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := parser.GetParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		service := parser.SchemaToEntityService(
			schema.Service{}, uint(id),
		)

		service, err = interactor.New(
			conf.Stor,
			conf.Logger,
			conf.Cache,
		).DeleteService(c, service)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, service)
	}
}
