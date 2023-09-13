package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

func (i interActor) CreatePass(
	ctx context.Context,
	pass entity.Password,
	userID uint,
) (entity.Password, error) {
	key, err := i.cache.GetPass(userID)
	if err != nil {
		return pass, err
	}
	return i.store.InsertPassword(ctx, pass, key)
}

func (i interActor) GetAllPassword(
	ctx context.Context,
	userID uint,
	decrypt bool,
) ([]entity.Password, error) {
	key, err := i.cache.GetPass(userID)
	if err != nil {
		return nil, err
	}
	return i.store.GetManyPassword(ctx, key, true)
}

func (i interActor) FindPassword(
	ctx context.Context,
	pass entity.Password,
	userID uint,
) (entity.Password, error) {
	key, err := i.cache.GetPass(userID)
	if err != nil {
		return pass, err
	}
	return i.store.GetPassword(ctx, pass, key, true)
}

func (i interActor) UpdatePass(
	ctx context.Context,
	pass entity.Password,
	userID uint,
) (entity.Password, error) {
	key, err := i.cache.GetPass(userID)
	if err != nil {
		return pass, err
	}
	return i.store.UpdatePassword(ctx, pass, key)
}

func (i interActor) DeletePass(
	ctx context.Context,
	pass entity.Password,
	userID uint,
) (entity.Password, error) {
	key, err := i.cache.GetPass(userID)
	if err != nil {
		return pass, err
	}
	return i.store.DeletePassword(ctx, pass, key)
}
