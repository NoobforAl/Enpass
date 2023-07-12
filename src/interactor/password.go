package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

func (i interActor) FindPassword(
	ctx context.Context,
	pass entity.Password,
	key string,
	decrypt bool,
) (entity.Password, error) {
	return i.store.GetPassword(
		ctx, pass, key, decrypt)
}

func (i interActor) GetAllPassword(
	ctx context.Context,
	key string,
	decrypt bool,
) ([]entity.Password, error) {
	return i.store.GetManyPassword(
		ctx, key, decrypt)
}

func (i interActor) UpdatePass(
	ctx context.Context,
	pass entity.Password,
	key string,
) (entity.Password, error) {
	return i.store.UpdatePassword(
		ctx, pass, key)
}

func (i interActor) CreatePass(
	ctx context.Context,
	pass entity.Password,
	key string,
) (entity.Password, error) {
	return i.store.InsertPassword(
		ctx, pass, key)
}

func (i interActor) DeletePass(
	ctx context.Context,
	pass entity.Password,
) (entity.Password, error) {
	return i.store.DeletePassword(
		ctx, pass)
}
