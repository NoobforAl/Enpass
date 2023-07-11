package schema

type service struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type CreateService struct{ service }
type UpdateService struct{ service }
