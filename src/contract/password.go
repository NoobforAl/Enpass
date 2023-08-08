package contract

import (
	"context"

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
		ctx context.Context,
		pass entity.Password,
		key string,
	) (entity.Password, error)
}

type ValidatePassword interface {
	ParsPassword(
		*gin.Context,
		*schema.Password,
	) error
}
