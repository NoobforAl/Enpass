package schema

type Service struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdateService struct {
	Service
	ServiceId uint `form:"serviceid" json:"serviceid" binding:"required"`
}
