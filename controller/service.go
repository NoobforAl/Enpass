package controller

import (
	"net/http"

	"github.com/NoobforAl/Enpass/Db"
	model "github.com/NoobforAl/Enpass/Model"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func NewService(c *gin.Context) {
	var ser schema.Service
	var err error

	if err = ser.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	service := model.Service{Name: model.Value(ser.Name)}
	if err = Db.Insert(&service); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, service)
}

func FindService(c *gin.Context) {
	serviceId, err := getParmInt(c, "id")

	if err != nil {
		errorHandling(c, err)
		return
	}

	service := model.Service{ID: uint(serviceId)}
	if err = Db.Get(&service); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, service)
}

func UpdateService(c *gin.Context) {
	var ser schema.UpdateService
	var err error

	if err = ser.Pars(c); err != nil {
		errorHandling(c, err)
		return
	}

	updateService := model.Service{ID: ser.ServiceId}
	if err = Db.Get(&updateService); err != nil {
		errorHandling(c, err)
		return
	}

	updateService.Name = model.Value(ser.Name)
	if err = Db.Update(&updateService); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, updateService)
}

func DeleteService(c *gin.Context) {
	serviceId, err := getParmInt(c, "id")
	if err != nil {
		errorHandling(c, err)
		return
	}

	err = Db.Delete(&model.Service{ID: uint(serviceId)})
	if err != nil {
		errorHandling(c, err)
		return
	}

	data, err := Db.GetMany(&model.SavedPassword{ServiceID: uint(serviceId)})
	if err != nil {
		errorHandling(c, err)
		return
	}

	if err = Db.DeleteMany(data); err != nil {
		errorHandling(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "record deleted"})
}
