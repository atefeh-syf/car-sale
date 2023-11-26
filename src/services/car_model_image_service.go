package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.CreateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.CreateCarModelImageRequest, dto.CarModelImageResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Image"},
			},
		},
	}
}

func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelImageService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelImageService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
