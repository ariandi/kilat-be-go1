package main

import (
	"database/sql"
	"gitlab.com/adl3905019/perumahan_go/api"
	"gitlab.com/adl3905019/perumahan_go/cache"
	util "gitlab.com/adl3905019/perumahan_go/utils"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config : ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	redis := cache.NewRedisClient(config)
	server, err := api.NewServer(config, store, redis)
	//server := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server", err)
	}
}
