package router

import (
	"github.com/NoobforAl/Enpass/http/v1/controller"
)

func Default(conf *controller.BaseConfig) {
	//r.Static("/", "./static/")

	conf.GinApp.POST("/login", controller.Login(conf))
	conf.GinApp.GET("/genRandomPass", controller.GenRandomPass)
}

func MainApi(conf *controller.BaseConfig) {
	api := conf.GinApp.Group("/api")
	api.Use(controller.AuthMiddleware(conf))

	api.PUT("/user", controller.UpdateUser(conf))

	api.GET("/password", controller.FindPass(conf))
	api.POST("/password", controller.NewPass(conf))
	api.GET("/password/:id", controller.AllPass(conf))
	api.PUT("/password/:id", controller.UpdatePass(conf))
	api.DELETE("/password/:id", controller.DeletePass(conf))

	api.GET("/service", controller.FindService(conf))
	api.POST("/service", controller.NewService(conf))
	api.GET("/service/:id", controller.AllService(conf))
	api.PUT("/service/:id", controller.UpdateService(conf))
	api.DELETE("/service/:id", controller.DeleteService(conf))
}
