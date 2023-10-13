package main

import (
	"log"

	"github.com/atefeh-syf/car-sale/api"
	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/data/cache"
	"github.com/atefeh-syf/car-sale/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InItRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}
	
	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}


	api.InitServer(cfg)
}