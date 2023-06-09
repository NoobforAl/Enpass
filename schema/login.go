package schema

type Login struct {
	Password string `form:"password" json:"password" binding:"required"`
}
