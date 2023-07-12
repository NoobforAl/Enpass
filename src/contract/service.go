package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
)

type Service interface {
	NewService(
		entity.Service,
	) (entity.Service, error)

	GetService(
		context.Context,
		entity.Service,
	) (entity.Service, error)

	GetManyService(
		context.Context,
	) ([]entity.Service, error)

	InsertService(
		context.Context,
		entity.Service,
	) (entity.Service, error)

	UpdateService(
		context.Context,
		entity.Service,
	) (entity.Service, error)

	DeleteService(
		context.Context,
		entity.Service,
	) (entity.Service, error)
}
