package config

import (
	"fmt"
	"log"
	"os"
)

type App struct {
	Api Api
	Database Database
	RedisConnectionString string
}

type Api struct{
	Address string
}

type Database struct{
	Username string
	Password string
	Host string
	DefaultDatabaseName string
}

func ReadFromEnvironmentVariables() App {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT variable must be set.")
	}

	cnf := App{
		Api: Api{Address: fmt.Sprintf(":%v", port)},
			Database: Database{
			Username:            os.Getenv("DB_USERNAME"),
			Password:            os.Getenv("DB_PASSWORD"),
			Host:                os.Getenv("DB_HOST"),
			DefaultDatabaseName: os.Getenv("DB_NAME"),
		},
		RedisConnectionString: os.Getenv("REDIS_URL"),
	}

	return cnf
}
