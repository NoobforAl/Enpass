package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
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

type ValidateService interface {
	ParsService(
		*gin.Context,
		*schema.Service,
	) error
}

type ParserService interface {
	SchemaToEntityService(
		ser schema.Service,
		id uint,
	) entity.Service

	EntityToDbModelService(
		ser entity.Service,
	) database.Service
}
