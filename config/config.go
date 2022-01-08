// Package config sets the application configurations using an external source.
package config

import (
	"fmt"
	"log"
	"os"
)

// App general application configuration variables.
type App struct {
	Api Api
	Database Database
	RedisConnectionString string
}

// Api represents api settings.
type Api struct{
	Address string
}

// Database represents database connection settings.
type Database struct{
	ConnectionString string
	DefaultDatabaseName string
}


// ReadFromEnvironmentVariables reads environment variables to set application configuration settings.
func ReadFromEnvironmentVariables() App {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT variable must be set.")
	}

	cnf := App{
		Api: Api{Address: fmt.Sprintf(":%v", port)},
		Database: Database{
			ConnectionString:    os.Getenv("DB_CONNECTION_STRING"),
			DefaultDatabaseName: os.Getenv("DB_NAME"),
		},
		RedisConnectionString: os.Getenv("REDIS_URL"),
	}

	return cnf
}
