package router

import (
	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/controller"
	"github.com/gin-gonic/gin"
)

func Default(r *gin.Engine, s contract.Stor) {
	r.GET("/genRandomPass", controller.GenRandomPass)
	r.POST("/login", controller.Login(s))
}

func MainApi(api *gin.RouterGroup, s contract.Stor) {
	api.Use(controller.AuthMiddleware())

	api.GET("/allPass", controller.AllPass(s))
	api.GET("/allService", controller.AllService(s))

	api.GET("/pass/:id", controller.FindPass(s))
	api.GET("/service/:id", controller.FindService(s))

	api.POST("/createPass", controller.NewPass(s))
	api.POST("/createService", controller.NewService(s))

	api.PUT("/updateUser", controller.UpdateUser(s))
	api.PUT("/updatePass", controller.UpdatePass(s))
	api.PUT("/updateService", controller.UpdateService(s))

	api.DELETE("/deletePass/:id", controller.DeletePassWord(s))
	api.DELETE("/deleteService/:id", controller.DeleteService(s))
}