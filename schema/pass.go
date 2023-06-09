package schema

type Pass struct {
	ServiceID uint   `form:"serviceid" json:"serviceid" binding:"required"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Note      string `form:"note" json:"note"`
}

type UpdatePass struct {
	Pass
	PassID uint `form:"passid" json:"passid" binding:"required"`
}
