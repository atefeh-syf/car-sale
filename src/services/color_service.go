package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type ColorService struct {
	base *BaseService[models.Color, dto.CreateColorRequest, dto.CreateColorRequest, dto.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dto.CreateColorRequest, dto.CreateColorRequest, dto.ColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

func (s *ColorService) Create(ctx context.Context, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *ColorService) Update(ctx context.Context, id int, req *dto.CreateColorRequest) (*dto.ColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *ColorService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *ColorService) GetById(ctx context.Context, id int) (*dto.ColorResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *ColorService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.ColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
