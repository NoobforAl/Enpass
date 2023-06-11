package contract

import (
	"context"

	"github.com/NoobforAl/Enpass/database"
)

type Service interface {
	NewService(id uint, name string) *database.Service

	GetService(ctx context.Context, id uint) (*database.Service, error)
	GetManyService(ctx context.Context) ([]*database.Service, error)

	InsertService(ctx context.Context, value database.Service) error
	InsertManyService(ctx context.Context, values []*database.Service) error

	UpdateService(ctx context.Context, m database.Service) error
	UpdateManyService(ctx context.Context, values []*database.Service) error

	DeleteService(ctx context.Context, id uint) error
	DeleteManyService(ctx context.Context, values []*database.Service) error
}

type User interface {
	NewUser(id uint, pass string) *database.User

	GetUser(ctx context.Context, id uint) (*database.User, error)
	GetManyUser(ctx context.Context) ([]*database.User, error)

	InsertUser(ctx context.Context, value database.User) error
	InsertManyUser(ctx context.Context, values []*database.User) error

	UpdateUser(ctx context.Context, m database.User) error
	UpdateManyUser(ctx context.Context, values []*database.User) error

	DeleteUser(ctx context.Context, id uint) error
	DeleteManyUser(ctx context.Context, values []*database.User) error
}

type Password interface {
	NewPassword(id, serId, userId uint,
		userName, password, note, hash string) *database.Password

	GetPassword(ctx context.Context, id uint) (*database.Password, error)
	GetManyPassword(ctx context.Context) ([]*database.Password, error)

	InsertPassword(ctx context.Context, value database.Password) error
	InsertManyPassword(ctx context.Context, values []*database.Password) error

	UpdatePassword(ctx context.Context, m database.Password) error
	UpdateManyPassword(ctx context.Context, values []*database.Password) error

	DeletePassword(ctx context.Context, id uint) error
	DeleteManyPassword(ctx context.Context, values []*database.Password) error
}

type Stor interface {
	User
	Service
	Password
}
