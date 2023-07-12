package parser

import (
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

func (_ schemaToEntity) Login(
	user schema.Login,
	id uint,
) entity.User {
	return entity.User{
		ID:       id,
		Password: user.Password,
	}
}

func (_ schemaToEntity) UpdateUser(
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

func (_ entityToDbModel) User(
	user entity.User,
) database.User {
	return database.User{
		ID:     user.ID,
		EnPass: user.Password,
	}
}
