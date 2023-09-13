package controller

import (
	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

type BaseConfig struct {
	GinApp     *gin.Engine
	Stor       contract.Store
	Logger     contract.Logger
	Cache      contract.Caching
	Validation contract.Validation
}
