package main

import (
	"github.com/atefeh-syf/car-sale/api"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/cache"
	"github.com/atefeh-syf/car-sale/data/db"
	"github.com/atefeh-syf/car-sale/data/db/migrations"
	"github.com/atefeh-syf/car-sale/pkg/logging"
)

// @securityDefinitions.apiKey AuthBearer
// @in header
// @name Authorization

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InItRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.StartUp, err.Error(), nil)
	}
	
	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)
}