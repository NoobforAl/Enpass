package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

type User interface {
	GetUser(
		context.Context,
		entity.User,
	) (entity.User, error)

	UpdateUser(
		ctx context.Context,
		old entity.User,
		new entity.User,
	) (entity.User, error)
}

type ValidateUser interface {
	ParsUpdateUser(
		*schema.UpdateUser,
	) error
}
