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
	l contract.Logger,
) {
	r.POST("/login",
		controller.Login(s, v, l))
	r.GET("/genRandomPass",
		controller.GenRandomPass)
}

func MainApi(
	r *gin.RouterGroup,
	s contract.Store,
	v contract.Validation,
	l contract.Logger,
) {
	r.Use(controller.AuthMiddleware(l))

	r.GET("/allPass",
		controller.AllPass(s, v, l))
	r.GET("/allService",
		controller.AllService(s, v, l))

	r.GET("/pass/:id",
		controller.FindPass(s, v, l))
	r.GET("/service/:id",
		controller.FindService(s, v, l))

	r.POST("/createPass",
		controller.NewPass(s, v, l))
	r.POST("/createService",
		controller.NewService(s, v, l))

	r.PUT("/updateUser/:id",
		controller.UpdateUser(s, v, l))
	r.PUT("/updatePass/:id",
		controller.UpdatePass(s, v, l))
	r.PUT("/updateService/:id",
		controller.UpdateService(s, v, l))

	r.DELETE("/deletePass/:id",
		controller.DeletePass(s, v, l))
	r.DELETE("/deleteService/:id",
		controller.DeleteService(s, v, l))
}
