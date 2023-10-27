package services

import (
	"fmt"
	"time"

	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/constants"
	"github.com/atefeh-syf/car-sale/data/cache"
	"github.com/atefeh-syf/car-sale/pkg/logging"
	"github.com/atefeh-syf/car-sale/pkg/service_errors"
	"github.com/go-redis/redis/v7"
)

type OtpService struct {
	logger logging.Logger
	cfg *config.Config
	redis *redis.Client
}

type OtpDto struct {
	Value string
	Used bool	
}

func NewOtpService(cfg *config.Config) * OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{ logger:  logger, cfg: cfg, redis: redis}
}



func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[OtpDto](s.redis, key)
	if err == nil && !res.Used{
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	} else if err == nil && res.Used{
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(s.redis, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redis, key)
	if err != nil {
		return err
	} else if err == nil && res.Used{
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if err == nil && !res.Used && res.Value != otp{
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else if err == nil && !res.Used && res.Value == otp{
		res.Used = true
		err = cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime * time.Second)
	}
	return nil
}
