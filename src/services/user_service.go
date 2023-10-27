package services

import (
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/common"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	cfg *config.Config
	logger logging.Logger
	database *gorm.DB
	otpService *OtpService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)

	return &UserService{
		cfg: cfg,
		logger:  logger,
		database: database,
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

