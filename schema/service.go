package schema

import "github.com/gin-gonic/gin"

type Service struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdateService struct {
	Service
	ServiceId uint `form:"serviceid" json:"serviceid" binding:"required"`
}

func (s *Service) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, s)
}

func (u *UpdateService) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, u)
}
