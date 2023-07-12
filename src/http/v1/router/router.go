package router

import (
	"github.com/NoobforAl/Enpass/contract"
	"github.com/NoobforAl/Enpass/http/v1/controller"
	"github.com/gin-gonic/gin"
)

func Default(
	r *gin.Engine,
	s contract.Store,
	v contract.Validation,
) {
	r.POST("/login",
		controller.Login(s, v))
	r.GET("/genRandomPass",
		controller.GenRandomPass)
}

func MainApi(
	api *gin.RouterGroup,
	s contract.Store,
	v contract.Validation,
) {
	api.Use(controller.AuthMiddleware())

	api.GET("/allPass",
		controller.AllPass(s, v))
	api.GET("/allService",
		controller.AllService(s, v))

	api.GET("/pass/:id",
		controller.FindPass(s, v))
	api.GET("/service/:id",
		controller.FindService(s, v))

	api.POST("/createPass",
		controller.NewPass(s, v))
	api.POST("/createService",
		controller.NewService(s, v))

	api.PUT("/updateUser",
		controller.UpdateUser(s, v))
	api.PUT("/updatePass",
		controller.UpdatePass(s, v))
	api.PUT("/updateService",
		controller.UpdateService(s, v))

	api.DELETE("/deletePass/:id",
		controller.DeletePass(s, v))
	api.DELETE("/deleteService/:id",
		controller.DeleteService(s, v))
}
