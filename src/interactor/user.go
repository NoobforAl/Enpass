package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

func (i interActor) FindUser(
	ctx context.Context,
	user schema.GetUser,
) (entity.User, error) {
	userPass := entity.User{
		ID:       1,
		Password: user.Password,
	}

	return i.store.GetUser(ctx, userPass)
}

func (i interActor) UpdateUser(
	ctx context.Context,
	user schema.UpdateUser,
) (entity.User, error) {

	oldPass := entity.User{
		Password: user.Old,
	}

	newPass := entity.User{
		Password: user.New,
	}

	return i.store.UpdateUser(
		ctx, oldPass, newPass,
	)
}
