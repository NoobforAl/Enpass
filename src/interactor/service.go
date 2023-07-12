package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

func (i interActor) FindService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	return i.store.GetService(
		ctx, ser)
}

func (i interActor) GetAllService(
	ctx context.Context,
) ([]entity.Service, error) {
	return i.store.GetManyService(ctx)
}

func (i interActor) CreateService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	return i.store.InsertService(
		ctx, ser)
}

func (i interActor) UpdateService(
	ctx context.Context,
	ser entity.Service,
	id uint,
) (entity.Service, error) {
	return i.store.UpdateService(
		ctx, ser)
}

func (i interActor) DeleteService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	return i.store.DeleteService(
		ctx, ser)
}
