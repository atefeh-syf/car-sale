package main

import (
	"github.com/atefeh-syf/car-sale/api"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InItRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}