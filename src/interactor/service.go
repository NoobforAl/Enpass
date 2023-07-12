package interactor

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

func (i interActor) FindService(
	ctx context.Context,
	id uint,
) (entity.Service, error) {
	return i.store.GetService(
		ctx, entity.Service{
			ServiceId: id,
		})
}

func (i interActor) GetAllService(
	ctx context.Context,
) ([]entity.Service, error) {
	return i.store.GetManyService(ctx)
}

func (i interActor) CreateService(
	ctx context.Context,
	service schema.CreateService,
) (entity.Service, error) {
	return i.store.InsertService(
		ctx, entity.Service{
			Name: service.Name,
		})
}

func (i interActor) UpdateService(
	ctx context.Context,
	upSer schema.UpdateService,
	id uint,
) (entity.Service, error) {
	return i.store.UpdateService(
		ctx, entity.Service{
			ServiceId: id,
			Name:      upSer.Name,
		})
}

func (i interActor) DeleteService(
	ctx context.Context,
	id uint,
) (entity.Service, error) {
	return i.store.DeleteService(
		ctx, entity.Service{
			ServiceId: id,
		})
}
