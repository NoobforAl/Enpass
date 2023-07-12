package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

func (i interActor) FindPassword(
	ctx context.Context,
	id uint,
	key string,
	decrypt bool,
) (entity.Password, error) {
	return i.store.GetPassword(
		ctx, entity.Password{
			PassID: id,
		}, key, decrypt)
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
	pass schema.UpdatePassword,
	id uint,
	key string,
) (entity.Password, error) {
	return i.store.UpdatePassword(
		ctx, entity.Password{
			PassID:    id,
			UserID:    1,
			ServiceID: pass.ServiceID,
			UserName:  pass.UserName,
			Password:  pass.Password,
			Note:      pass.Note,
		})
}

func (i interActor) CreatePass(
	ctx context.Context,
	pass schema.UpdatePassword,
	key string,
) (entity.Password, error) {
	return i.store.InsertPassword(
		ctx, entity.Password{
			UserID:    1,
			ServiceID: pass.ServiceID,
			UserName:  pass.UserName,
			Password:  pass.Password,
			Note:      pass.Note,
		}, key)
}

func (i interActor) DeletePass(
	ctx context.Context,
	id uint,
) (entity.Password, error) {
	return i.store.DeletePassword(
		ctx, entity.Password{
			PassID: id,
		})
}
