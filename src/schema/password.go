package schema

type Password struct {
	ServiceID uint   `form:"serviceid" json:"serviceid" binding:"required"`
	UserName  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Note      string `form:"note" json:"note"`
}
