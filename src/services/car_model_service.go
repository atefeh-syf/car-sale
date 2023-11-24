package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type CarModelService struct {
	base *BaseService[models.CarModel, dto.CreateCarModelRequest, dto.CreateCarModelRequest, dto.CarModelResponse]
}

func NewCarModelService(cfg *config.Config) *CarModelService {
	return &CarModelService{
		base: &BaseService[models.CarModel, dto.CreateCarModelRequest, dto.CreateCarModelRequest, dto.CarModelResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Company.Country"},
				{string: "CarType"},
				{string: "Gearbox"},
				{string: "CarModelColors.Color"},
				{string: "CarModelYears.PersianYear"},
				{string: "CarModelYears.CarModelPriceHistories"},
				{string: "CarModelProperties.Property.Category"},
				{string: "CarModelImages.Image"},
				{string: "CarModelComments.User"},
			},
		},
	}
}

func (s *CarModelService) Create(ctx context.Context, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CarModelService) Update(ctx context.Context, id int, req *dto.CreateCarModelRequest) (*dto.CarModelResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CarModelService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CarModelService) GetById(ctx context.Context, id int) (*dto.CarModelResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CarModelService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
