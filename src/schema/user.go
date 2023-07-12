package schema

type UpdateUser struct {
	Old string `form:"old" json:"old" binding:"required"`
	New string `form:"new" json:"new" binding:"required"`
}

type Login struct {
	Password string `form:"password" json:"password" binding:"required"`
}
