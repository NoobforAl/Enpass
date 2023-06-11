package entity

import (
	"context"

	"github.com/NoobforAl/Enpass/contract"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ServiceId uint   `form:"serviceid" json:"serviceid"`
	Name      string `form:"name" json:"name"`
}

func (s *Service) Pars(c *gin.Context) error {
	return parsJsonAndValidate(c, s)
}

func (s *Service) CreateService(ctx context.Context, stor contract.Stor) error {
	ser := stor.NewService(s.ServiceId, s.Name)
	return stor.InsertService(ctx, *ser)
}

func (s *Service) UpdateService(ctx context.Context, stor contract.Stor) error {
	ser, err := stor.GetService(ctx, s.ServiceId)
	if err != nil {
		return err
	}

	ser = stor.NewService(ser.ID, s.Name)
	return stor.InsertService(ctx, *ser)
}

func (s *Service) DeleteService(ctx context.Context, stor contract.Stor) error {
	var err error
	if err = stor.DeleteService(ctx, s.ServiceId); err != nil {
		return err
	}

	passwords, err := stor.GetManyPassword(ctx)
	if err != nil {
		return err
	}

	for _, v := range passwords {
		if v.ServiceID == s.ServiceId {
			if err = stor.DeletePassword(ctx, v.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Service) FindService(ctx context.Context, stor contract.Stor) (*Service, error) {
	ser, err := stor.GetService(ctx, s.ServiceId)
	if err != nil {
		return nil, err
	}
	return &Service{ServiceId: s.ServiceId, Name: ser.Name.String()}, nil
}

func (s *Service) GetAllService(ctx context.Context, stor contract.Stor) ([]Service, error) {
	services, err := stor.GetManyService(ctx)
	if err != nil {
		return nil, err
	}

	allService := make([]Service, len(services))
	for i := range allService {
		allService[i].Name = services[i].Name.String()
		allService[i].ServiceId = services[i].ID
	}
	return allService, nil
}
