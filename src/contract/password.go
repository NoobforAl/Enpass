package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

type Password interface {
	GetPassword(
		ctx context.Context,
		pass entity.Password,
		key string,
		decrypt bool,
	) (entity.Password, error)

	GetManyPassword(
		ctx context.Context,
		key string,
		decrypt bool,
	) ([]entity.Password, error)

	InsertPassword(
		ctx context.Context,
		pass entity.Password,
		key string,
	) (entity.Password, error)

	UpdatePassword(
		ctx context.Context,
		pass entity.Password,
		key string,
	) (entity.Password, error)

	DeletePassword(
		context.Context,
		entity.Password,
	) (entity.Password, error)
}

type ValidatePassword interface {
	ParsPassword(
		*gin.Context,
		*schema.Password,
	) error
}

type ParserPassword interface {
	SchemaToEntityPass(
		pass schema.Password,
		passID, userId uint,
	) entity.Password

	EntityToDbModelPass(
		pass entity.Password,
	) database.Password
}
