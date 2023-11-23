package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Category"},
			},
		},
	}
}

func (s *PropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyService) Update(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *PropertyService) GetById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *PropertyService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
