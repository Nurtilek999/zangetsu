package main

import (
	"log"
	"zangetsu/api"
	"zangetsu/pkg/config"
	"zangetsu/pkg/database"
)

func init() {
	config.GetConfig()
}

func main() {
	pgdb, err := database.InitDB()
	if err != nil {
		log.Fatal("error in Postgres connection: ", err.Error())
		return
	}
	defer pgdb.Close()

	esdb, err := database.InitESDb()
	if err != nil {
		log.Fatal("error in ElasticSearch connection:", err.Error())
		return
	}
	defer esdb.Stop()

	port := ":8000"
	app := api.SetupRouter(pgdb, esdb)
	app.Run(port)
}
