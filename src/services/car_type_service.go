package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dto.CreateCarTypeRequest, dto.CreateCarTypeRequest, dto.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dto.CreateCarTypeRequest, dto.CreateCarTypeRequest, dto.CarTypeResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "CarModels"},
			},
		},
	}
}

func (s *CarTypeService) Create(ctx context.Context, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarTypeService) Update(ctx context.Context, id int, req *dto.CreateCarTypeRequest) (*dto.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarTypeService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarTypeService) GetById(ctx context.Context, id int) (*dto.CarTypeResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarTypeService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
