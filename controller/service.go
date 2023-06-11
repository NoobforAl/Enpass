package controller

import (
	"net/http"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/gin-gonic/gin"
)

func NewService(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		var err error

		if err = ser.Pars(c); err != nil {
			errorHandling(c, err)
			return
		}

		if err = ser.CreateService(c, stor); err != nil {
			errorHandling(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func AllService(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		services, err := ser.GetAllService(c, stor)
		if err != nil {
			errorHandling(c, err)
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

func FindService(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceId, err := getParmInt(c, "id")
		if err != nil {
			errorHandling(c, err)
			return
		}

		ser := &entity.Service{ServiceId: uint(serviceId)}
		ser, err = ser.FindService(c, stor)
		if err != nil {
			errorHandling(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func UpdateService(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ser entity.Service
		var err error

		if err = ser.Pars(c); err != nil {
			errorHandling(c, err)
			return
		}

		if err = ser.UpdateService(c, stor); err != nil {
			errorHandling(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}

func DeleteService(stor contract.Stor) gin.HandlerFunc {
	return func(c *gin.Context) {
		serviceId, err := getParmInt(c, "id")
		if err != nil {
			errorHandling(c, err)
			return
		}

		ser := entity.Service{ServiceId: uint(serviceId)}
		if err = ser.DeleteService(c, stor); err != nil {
			errorHandling(c, err)
			return
		}

		c.JSON(http.StatusOK, ser)
	}
}
