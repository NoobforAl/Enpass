package validation

import (
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

func ParsUser(
	c *gin.Context,
	user *schema.UpdateUser,
) error {
	return c.BindJSON(user)
}
