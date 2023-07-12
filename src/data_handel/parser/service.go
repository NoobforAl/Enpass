package parser

import (
	"github.com/NoobforAl/Enpass/database"
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

// schema to entity service
func (_ parser) SchemaToEntityService(
	ser schema.Service,
	id uint,
) entity.Service {
	return entity.Service{
		ServiceId: id,
		Name:      ser.Name,
	}
}

func (_ parser) EntityToDbModelService(
	ser entity.Service,
) database.Service {
	return database.Service{
		ID:   ser.ServiceId,
		Name: ser.Name,
	}
}
