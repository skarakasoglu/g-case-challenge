package config

import "os"

type App struct {
	Database Database
}

type Database struct{
	Username string
	Password string
	Host string
	DefaultDatabaseName string
}

func ReadFromEnvironmentVariables() App {
	cnf := App{
		Database: Database{
		Username:            os.Getenv("DB_USERNAME"),
		Password:            os.Getenv("DB_PASSWORD"),
		Host:                os.Getenv("DB_HOST"),
		DefaultDatabaseName: os.Getenv("DB_NAME"),
	}}

	return cnf
}
