package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: &BaseService[models.PropertyCategory, dto.CreatePropertyCategoryRequest, dto.UpdatePropertyCategoryRequest, dto.PropertyCategoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Properties"},
			},
		},
	}
}

func (s *PropertyCategoryService) Create(ctx context.Context, req *dto.CreatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *PropertyCategoryService) Update(ctx context.Context, id int, req *dto.UpdatePropertyCategoryRequest) (*dto.PropertyCategoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *PropertyCategoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *PropertyCategoryService) GetById(ctx context.Context, id int) (*dto.PropertyCategoryResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *PropertyCategoryService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyCategoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
