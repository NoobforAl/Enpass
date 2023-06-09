package schema

type UpdateUserPass struct {
	Password string `form:"password" json:"password" binding:"required"`
}
