package services

import (
	"context"
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.CreateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyRequest, dto.CreateCompanyRequest, dto.CompanyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "CarModels"},
			},
		},
	}
}

func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *CompanyService) Update(ctx context.Context, id int, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Update(ctx, id, req)
}

func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetById(ctx, id)
}

func (s *CompanyService) GetByFiler(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
