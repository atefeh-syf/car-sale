package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.CreateCarModelColorRequest, dto.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dto.CreateCarModelColorRequest, dto.CreateCarModelColorRequest, dto.CarModelColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Color"},
			},
		},
	}
}

func (s *CarModelColorService) Create(ctx context.Context, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelColorService) Update(ctx context.Context, id int, req *dto.CreateCarModelColorRequest) (*dto.CarModelColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelColorService) GetById(ctx context.Context, id int) (*dto.CarModelColorResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelColorService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
