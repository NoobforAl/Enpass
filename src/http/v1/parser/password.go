package parser

import (
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

// schema to entity password
func SchemaToEntityPass(
	pass schema.Password,
	passID, userId uint,
) entity.Password {
	return entity.Password{
		ID:        passID,
		UserID:    userId,
		ServiceID: pass.ServiceID,
		UserName:  pass.UserName,
		Password:  pass.Password,
		Note:      pass.Note,
	}
}
