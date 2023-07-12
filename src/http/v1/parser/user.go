package parser

import (
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

func SchemaToEntityLogin(
	user schema.Login,
	id uint,
) entity.User {
	return entity.User{
		ID:       id,
		Password: user.Password,
	}
}

func SchemaToEntityUser(
	user schema.UpdateUser,
	id uint,
) (old, new entity.User) {
	return entity.User{
			ID:       id,
			Password: user.Old,
		},
		entity.User{
			ID:       id,
			Password: user.New,
		}
}
