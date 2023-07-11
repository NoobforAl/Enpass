package schema

type password struct {
	ServiceID uint   `form:"serviceid" json:"serviceid" binding:"required"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Note      string `form:"note" json:"note"`
}

type CreatePassword struct{ password }
type UpdatePassword struct{ password }
