package services

import (
	"github.com/atefeh-syf/car-sale/api/dto"
	"github.com/atefeh-syf/car-sale/common"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/constants"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/models"
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

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error;err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByMobileNumber(MobileNumber string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).Select("count(*) > 0 ").Where("mobile_number = ?", MobileNumber).Find(&exists).Error;err != nil{
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return true, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {
	if err := s.database.Model(&models.Role{}).Select("id").Where("name = ?", constants.DefaultRoleName).Find(&roleId).Error;err != nil {
		return 0, err
	}
	return roleId, nil
}
