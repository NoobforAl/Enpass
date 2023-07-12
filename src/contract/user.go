package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/entity"
	"github.com/NoobforAl/Enpass/schema"
	"github.com/gin-gonic/gin"
)

type User interface {
	GetUser(
		context.Context,
		entity.User,
	) (entity.User, error)

	UpdateUser(
		ctx context.Context,
		old entity.User,
		new entity.User,
	) (entity.User, error)
}

type ValidateUser interface {
	ParsUpdateUser(
		*gin.Context,
		*schema.UpdateUser,
	) error

	ParsLoginUser(
		*gin.Context,
		*schema.Login,
	) error
}
