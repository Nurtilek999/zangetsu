package main

import (
	"zangetsu/api"
	"zangetsu/pkg/config"
	"zangetsu/pkg/database"
	"zangetsu/pkg/logging"
)

func init() {
	config.GetConfig()
}

func main() {
	logger := logging.GetLogger()
	pgdb, err := database.InitDB()
	if err != nil {
		//log.Fatal("error in Postgres connection: ", err.Error())
		logger.Error("error in Postgres connection")
		return
	}
	defer pgdb.Close()

	esdb, err := database.InitESDb()
	if err != nil {
		//log.Fatal("error in ElasticSearch connection:", err.Error())
		logger.Error("error in ElasticSearch connection")
		return
	}
	defer esdb.Stop()

	port := ":8000"
	app := api.SetupRouter(pgdb, esdb, logger)
	app.Run(port)
}
