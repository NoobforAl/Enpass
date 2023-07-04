package database

import (
	"context"
	"time"
)

type Service struct {
	ID        uint  `gorm:"primarykey;uniq"`
	Name      Value `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Stor) NewService(id uint, name string) *Service {
	return &Service{ID: id, Name: Value(name)}
}

func (s Stor) GetService(ctx context.Context, id uint) (*Service, error) {
	ser := s.NewService(id, "")
	return ser, get(ctx, ser)
}

func (s Stor) GetManyService(ctx context.Context) ([]*Service, error) {
	return getMany(ctx, s.NewService(0, ""))
}

func (s Stor) InsertService(ctx context.Context, value Service) error {
	return insert(ctx, &value)
}

func (s Stor) InsertManyService(ctx context.Context, values []*Service) error {
	return insertMany(ctx, values)
}

func (s Stor) UpdateService(ctx context.Context, m Service) error {
	err := get(ctx, s.NewService(m.ID, ""))
	if err != nil {
		return err
	}
	return update(ctx, &m)
}

func (s Stor) UpdateManyService(ctx context.Context, values []*Service) error {
	return updateMany(ctx, values)
}

func (s Stor) DeleteService(ctx context.Context, id uint) error {
	return delete(ctx, s.NewService(id, ""))
}

func (s Stor) DeleteManyService(ctx context.Context, values []*Service) error {
	return deleteMany(ctx, values)
}
