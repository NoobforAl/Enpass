package database

import (
	"context"
	"time"

	"github.com/NoobforAl/Enpass/entity"
)

type Service struct {
	ID        uint   `gorm:"primarykey;uniq"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func entityToModelService(
	ser entity.Service,
) Service {
	return Service{
		ID:   ser.ID,
		Name: ser.Name,
	}
}

func modelToEntityService(
	ser Service,
) entity.Service {
	return entity.Service{
		ID:   ser.ID,
		Name: ser.Name,
	}
}

func (s Stor) GetService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	service := entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		First(&service).Error

	return modelToEntityService(service), err
}

func (s Stor) GetManyService(
	ctx context.Context,
) ([]entity.Service, error) {
	var data []*Service

	err := s.db.Model(&Service{}).
		WithContext(ctx).
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	services := make([]entity.Service, len(data))
	for i := range services {
		services[i] = modelToEntityService(*data[i])
	}
	return services, nil
}

func (s Stor) InsertService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	service := entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Save(&service).Error

	return modelToEntityService(service), err
}

func (s Stor) UpdateService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	service := entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		Save(&service).Error

	return modelToEntityService(service), err
}

func (s Stor) DeleteService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	service := entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		First(&service).
		Delete(&service).Error

	return modelToEntityService(service), err
}
