package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.CreateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.CreateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *GearboxService) Update(ctx context.Context, id int, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *GearboxService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
