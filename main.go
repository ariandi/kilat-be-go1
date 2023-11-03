package main

import (
	"github.com/ariandi/kilat-be-go1/api"
	"github.com/ariandi/kilat-be-go1/cache"
	"github.com/ariandi/kilat-be-go1/db"
	util "github.com/ariandi/kilat-be-go1/utils"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config : ", err)
	}
	config.PdamCd = util.LoadPdamCd()
	config.PdamAdmin = util.LoadPdamAdmin()
	config.TrxConstant = util.LoadTrxConstant()

	mongoDb, err := db.Connect(config)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := mongoDb.Database(config.MongoDbName)
	redis := cache.NewRedisClient(config)
	server, err := api.NewServer(config, redis, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server", err)
	}
}
