package db

import (
	"fmt"
	"time"

	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(cfg *config.Config)  error{
	var err error
	cnn := fmt.Sprintf("host=%s port=%s  user=%s  password=%s dbname=%s sslmode=%s TimeZone=Asia/tehran", 
											cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)
	dbClient, err = gorm.Open(postgres.Open(cnn),  &gorm.Config{})

	if err != nil{
		return err
	}

	sqlDb , _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil{
		return err
	}
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.StartUp, "DB Connection  established", nil)
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb() {
	con , _ := dbClient.DB()
	con.Close()
}

