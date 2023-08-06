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

func (s Stor) entityToModelService(ser entity.Service) Service {
	s.log.Debug("Entity To Model Service")
	return Service{
		ID:   ser.ID,
		Name: ser.Name,
	}
}

func (s Stor) modelToEntityService(ser Service) entity.Service {
	s.log.Debug("Model To Entity Service")
	return entity.Service{
		ID:   ser.ID,
		Name: ser.Name,
	}
}

func (s Stor) GetService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	s.log.Debug("Get Service")
	service := s.entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		First(&service).Error

	return s.modelToEntityService(service), err
}

func (s Stor) GetManyService(
	ctx context.Context,
) ([]entity.Service, error) {
	s.log.Debug("Get Many Service")
	var data []*Service

	if err := s.db.Model(&Service{}).
		WithContext(ctx).
		Find(&data).Error; err != nil {
		return nil, err
	}

	services := make([]entity.Service, len(data))
	for i := range services {
		services[i] = s.modelToEntityService(*data[i])
	}
	return services, nil
}

func (s Stor) InsertService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	s.log.Debug("Insert Service")
	service := s.entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Save(&service).Error

	return s.modelToEntityService(service), err
}

func (s Stor) UpdateService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	s.log.Debug("Update Service")
	service := s.entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		Save(&service).Error

	return s.modelToEntityService(service), err
}

func (s Stor) DeleteService(
	ctx context.Context,
	ser entity.Service,
) (entity.Service, error) {
	s.log.Debug("Delete Service")
	service := s.entityToModelService(ser)

	err := s.db.Model(&service).
		WithContext(ctx).
		Where("id = ?", service.ID).
		First(&service).
		Delete(&service).Error

	return s.modelToEntityService(service), err
}
