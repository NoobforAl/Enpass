package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

func (i interActor) FindUser(
	ctx context.Context,
	user entity.User,
) (entity.User, error) {
	return i.store.GetUser(ctx, user)
}

func (i interActor) UpdateUser(
	ctx context.Context,
	oldUser entity.User,
	newUser entity.User,
) (entity.User, error) {
	return i.store.UpdateUser(
		ctx, oldUser, newUser,
	)
}
