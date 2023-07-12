package parser

import (
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

// schema to entity password
func (_ parser) SchemaToEntityPass(
	pass schema.Password,
	passID, userId uint,
) entity.Password {
	return entity.Password{
		PassID:    passID,
		UserID:    userId,
		ServiceID: pass.ServiceID,
		UserName:  pass.UserName,
		Password:  pass.Password,
		Note:      pass.Note,
	}
}

func (_ parser) EntityToDbModelPass(
	pass entity.Password,
) database.Password {
	return database.Password{
		ID:        pass.PassID,
		UserID:    pass.UserID,
		ServiceID: pass.ServiceID,
		Values: database.Values{
			UserName: pass.UserName,
			Password: pass.Password,
			Note:     pass.Note,
		},
	}
}
