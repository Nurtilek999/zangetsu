package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func InitDB() (*sql.DB, error) {
	username := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	dbname := viper.GetString("Database.DBName")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("error in db connection: ", err.Error())
	}

	return DB, nil
}
