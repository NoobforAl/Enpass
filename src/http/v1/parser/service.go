package parser

import (
	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
)

// schema to entity service
func SchemaToEntityService(
	ser schema.Service,
	id uint,
) entity.Service {
	return entity.Service{
		ID:   id,
		Name: ser.Name,
	}
}
