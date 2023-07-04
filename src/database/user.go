package database

import (
	"context"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	EnPass    Value
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Stor) NewUser(id uint, pass string) *User {
	return &User{ID: id, EnPass: Value(pass)}
}

func (s Stor) GetUser(ctx context.Context, id uint) (*User, error) {
	ser := s.NewUser(id, "")
	return ser, get(ctx, ser)
}

func (s Stor) GetManyUser(ctx context.Context) ([]*User, error) {
	return getMany(ctx, s.NewUser(0, ""))
}

func (s Stor) InsertUser(ctx context.Context, value User) error {
	return insert(ctx, &value)
}

func (s Stor) InsertManyUser(ctx context.Context, values []*User) error {
	return insertMany(ctx, values)
}

func (s Stor) UpdateUser(ctx context.Context, m User) error {
	err := get(ctx, s.NewUser(m.ID, ""))
	if err != nil {
		return err
	}
	return update(ctx, &m)
}

func (s Stor) UpdateManyUser(ctx context.Context, values []*User) error {
	return updateMany(ctx, values)
}

func (s Stor) DeleteUser(ctx context.Context, id uint) error {
	return delete(ctx, s.NewUser(id, ""))
}

func (s Stor) DeleteManyUser(ctx context.Context, values []*User) error {
	return deleteMany(ctx, values)
}
