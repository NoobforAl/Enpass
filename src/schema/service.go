package schema

type Service struct {
	Name string `form:"name" json:"name" binding:"required"`
}
