package controller

import (
	"net/http"

	errs "github.com/NoobforAl/Enpass/errors"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/gin-gonic/gin"
)

func NewService(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		var err error

		if err = ser.Pars(c); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		if err = ser.CreateService(c, stor); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func AllService(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		services, err := ser.GetAllService(c, stor)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

func FindService(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceId, err := getParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		ser := &entity.Service{ServiceId: uint(serviceId)}
		ser, err = ser.FindService(c, stor)
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func UpdateService(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		var err error

		if err = ser.Pars(c); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		if err = ser.UpdateService(c, stor); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func DeleteService(stor contract.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceId, err := getParmInt(c, "id")
		if err != nil {
			errs.ErrHandle(c, err)
			return
		}

		ser := entity.Service{ServiceId: uint(serviceId)}
		if err = ser.DeleteService(c, stor); err != nil {
			errs.ErrHandle(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}
