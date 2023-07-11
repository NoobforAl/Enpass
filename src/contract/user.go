package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
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
