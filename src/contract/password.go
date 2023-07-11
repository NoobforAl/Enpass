package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

type Password interface {
	NewPassword(
		entity.Password,
	) (entity.Password, error)

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
		context.Context,
		entity.Password,
	) (entity.Password, error)

	UpdatePassword(
		context.Context,
		entity.Password,
	) (entity.Password, error)

	DeletePassword(
		context.Context,
		entity.Password,
	) (entity.Password, error)
}
