package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
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
		ctx context.Context,
		pass entity.Password,
		key string,
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

type ValidatePassword interface {
	ParsCreatePassword(
		*schema.CreatePassword,
	) error

	ParsUpdatePassword(
		*schema.UpdatePassword,
	) error
}
